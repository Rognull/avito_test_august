package db

import (
	"avito_test/internals/models"
	"l0/internal/model"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"context"

)

type Storage struct {
	databasePool *pgxpool.Pool
}



func NewStorage(dBase *pgxpool.Pool) *Storage {
	storage := new(Storage)
	storage.databasePool = dBase
	return storage
} 

func (db *Storage) AddSegment(order model.Order) error {
 
}

func (db *Storage) AddUserSegment(order model.Order) error {
 
}
 

func (db *Storage) DeleteSegment(order model.Order) error {
 
}
 
func (db *Storage) FindUserSegment(slug string) ([]models.UserSegment,error) {
	query := "SELECT segments.id, segments.slug,segments.created_at FROM segments JOIN user_segments ON user_segments.segment_id = segments.id  WHERE user_id = $1" // TODO right query 
	 
	var result []models.UserSegment

	err := pgxscan.Select(context.Background(), db.databasePool, &result, query, slug)

	if err != nil {
		return nil, fmt.Errorf("FindUserSegment: failed to execute query: %w", err)
	}

	return result,err

}
 