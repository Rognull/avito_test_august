package db

import (
	"l0/internal/model"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
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
 
func (db *Storage) FindUserSegment(order model.Order) error {
 
}
 