package db

import (
	"avito_test/internals/models"
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

type Storage struct {
	databasePool *pgxpool.Pool
}

func NewStorage(dBase *pgxpool.Pool) *Storage {
	storage := new(Storage)
	storage.databasePool = dBase
	return storage
}

func (db *Storage) AddSegment(seg models.Segment) ([]models.User, error) {

	ctx := context.Background()
	tx, err := db.databasePool.Begin(ctx)
	defer func() {
		err = tx.Rollback(context.Background())
		if err != nil {
			log.Errorln(err)
			// fmt.Printf("Error in Rollback: %w", err)
		}
	}()
    queryy := "CREATE TABLE IF NOT EXISTS segments ( id SERIAL PRIMARY KEY, slug VARCHAR NOT NULL UNIQUE, created_at TIMESTAMP DEFAULT current_timestamp)" 
	db.databasePool.Exec(context.Background(),queryy)
	queryForSegment := "INSERT INTO segments (id, slug, created_at) VALUES (DEFAULT, $1, DEFAULT) RETURNING id"

	idOfSegment:= -1
	err = tx.QueryRow(context.Background(),queryForSegment,seg.Slug).Scan(&idOfSegment)
	
	// Exec(ctx, queryForSegment, seg.Slug,&idOfSegment)

	if err != nil {
		log.Errorln(err)
		return nil, fmt.Errorf("Error in AddSegment, at Insert into table: %W", err)
	}

	buf := []models.User{}
	if seg.Procent != 0 {
		 
		procent := seg.Procent / 100.0
		query := "SELECT id FROM users LIMIT(SELECT COUNT(*) FROM my_table) * $1 "
		err = pgxscan.Get(ctx, tx, buf, query, procent)

		if err != nil {
			log.Errorln(err)
			err = tx.Rollback(context.Background()) 
			if err != nil {
				log.Errorln(err)
			}
			return nil, fmt.Errorf("error in select query: %w", err)
		}

		query = "INSERT INTO user_segments(user_id,segment_id,added_at,delete_time) VALUES ($1,$2,DEFAULT,DEFAULT);"

		for _,i:=range buf{
		 _,err = tx.Exec(ctx, queryForSegment,i.ID,idOfSegment)
		 if err != nil {
			log.Errorln(err)
			err = tx.Rollback(context.Background())  
			if err != nil {
				log.Errorln(err)
			}
			return nil, fmt.Errorf("error in select query: %w", err)
		 }
		}
	}
	err = tx.Commit(context.Background())
	if err != nil {
		log.Errorln(err)
	}
	return buf, err
}


func (db *Storage) AddUserSegment(order models.UserSegment) error {
return nil
}

func (db *Storage) DeleteSegment(order models.Segment) error {
return nil
}

func (db *Storage) FindUserSegment(slug string) ([]models.UserSegment, error) {
	query := "SELECT segments.id, segments.slug,segments.created_at FROM segments JOIN user_segments ON user_segments.segment_id = segments.id  WHERE user_id = $1" // TODO right query

	var result []models.UserSegment

	err := pgxscan.Select(context.Background(), db.databasePool, &result, query, slug)

	if err != nil {
		return nil, fmt.Errorf("FindUserSegment: failed to execute query: %w", err)
	}

	return result, err

}
