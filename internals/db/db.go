package db

import (
	"avito_test/internals/models"
	"context"
	"fmt"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type storage struct {
	databasePool *pgxpool.Pool
}


 type Storage interface{
	AddSegment(models.Segment) ([]models.User, error)
	DeleteSegment(string) error
	FindUserSegment(int64) ([]models.Segment, error) 
	AddUserSegment( models.AddRequest) error
 }

func NewStorage(dBase *pgxpool.Pool) *storage {
	storage := new(storage)
	storage.databasePool = dBase
	return storage
}

func (db *storage) AddSegment(seg models.Segment) ([]models.User, error) {

	ctx := context.Background()
	tx, err := db.databasePool.Begin(ctx)
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			// log.Errorln(err)
			// fmt.Printf("Error in Rollback: %w", err)
		}
	}()
    // queryy := "CREATE TABLE IF NOT EXISTS segments ( id SERIAL PRIMARY KEY, slug VARCHAR NOT NULL UNIQUE, created_at TIMESTAMP DEFAULT current_timestamp)" 
	// db.databasePool.Exec(context.Background(),queryy)
	queryForSegment := "INSERT INTO segments (id, slug, created_at) VALUES (DEFAULT, $1, DEFAULT) RETURNING id"

	idOfSegment := -1
	err = tx.QueryRow(context.Background(),queryForSegment,seg.Slug).Scan(&idOfSegment)
	// Exec(ctx, queryForSegment, seg.Slug,&idOfSegment)

	if err != nil {
		log.Errorln(err)
		return nil, fmt.Errorf("Error in AddSegment, at Insert into table: %W", err)
	}

	buf := []models.User{}
	if seg.Procent != 0 {
		procent := float64(seg.Procent) / 100.0
		// query := "SELECT id FROM users LIMIT(SELECT COUNT(id) FROM users) * $1 "
		query := "SELECT id FROM users"
		err = pgxscan.Select(context.Background(), tx, &buf, query)
		fmt.Println(buf)
		if err != nil {
			log.Errorln(err)
			err = tx.Rollback(context.Background()) 
			if err != nil {
				log.Errorln(err)
			}
			return nil, fmt.Errorf("error in select query: %w", err)
		}

		query = "INSERT INTO user_segments(user_id, segment_id, added_at, delete_time) VALUES ($1, $2, DEFAULT, DEFAULT);"
		 
		for i:=0; i < int(float64(len(buf)) * procent);i++{
			log.Println(i, buf[i].ID)
		 _,err = tx.Exec(context.Background(), query ,buf[i].ID,idOfSegment)
		 if err != nil {
			log.Errorln(err)
			err = tx.Rollback(context.Background())  
			if err != nil {
				log.Errorln(err)
			}
			return nil, fmt.Errorf("error in insert query: %w", err)
		 }
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Errorln(err)
	}
	return buf, err
}


func (db *storage) AddUserSegment(addRequest models.AddRequest) error {
	
	query := "INSERT INTO user_segments(user_id, segment_id) (SELECT $1, id from segments where slug = ANY($2))"

   _, err := db.databasePool.Exec(context.Background(),query,addRequest.ID, pq.Array(addRequest.AddSegment))

		if err != nil{ 
			return   fmt.Errorf("AddUserSegment: failed to execute query: %w", err)
		}
	return nil 
}

// func (db *storage) DeleteUserSegment(models.AddRequest) error {
	
// 	query:=  "UPDATE user_segments SET delete_time = CURRENT_TIMESTAMP::timestamp  where segment_id = (SELECT id from segments where slug = $1)"
	
// }

// func (db *storage) GetSegmentId(slug []string) ([]models.Segment, error){

// 	query:=  "SELECT * FROM segments where slug = $1 "

// 	var result []models.Segment

// 	for _,i := range slug {
// 		err := pgxscan.Select(context.Background(), db.databasePool, &result, query, i)
// 		if err != nil{ 
// 			return nil, fmt.Errorf("DeleteSegment: failed to execute query: %w", err)
// 		}

// 	}
// 	return result, nil
// }

func (db *storage) DeleteSegment(slug string) error {

	query:=  "UPDATE user_segments SET delete_time = CURRENT_TIMESTAMP::timestamp  where segment_id = (SELECT id from segments where slug = $1)"
	 
	_ , err := db.databasePool.Exec(context.Background(),query,slug)

	if err != nil {
		return fmt.Errorf("DeleteSegment: failed to execute query: %w", err)
	}

	return err
}

func (db *storage) FindUserSegment(id int64) ([]models.Segment, error) {
	query := "SELECT segments.id, segments.slug, segments.created_at FROM segments JOIN user_segments ON user_segments.segment_id = segments.id  WHERE user_id = $1 AND (delete_time > CURRENT_TIMESTAMP::timestamp or delete_time is NULL)" // TODO right query

	var result []models.Segment

	err := pgxscan.Select(context.Background(), db.databasePool, &result, query, id)

	if err != nil {
		return nil, fmt.Errorf("FindUserSegment: failed to execute query: %w", err)
	}

	return result, err

}
