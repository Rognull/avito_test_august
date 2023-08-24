package services

import (
	"errors"
 
	// "fmt"
	"avito_test/internals/db"
	"avito_test/internals/models"
	"github.com/sirupsen/logrus"
)

type Service struct{
	Storage db.Storage
}


func NewService(storage *db.Storage) *Service{
	resultService := new(Service)
	resultService.Storage = *storage
	return resultService
}



func (s *Service) FindUserSegment(Slug string) (models.User ,error){    
 
}

func (s *Service) NewUserSegment(segment models.UserSegment) (error){     
 
}
 
func (s *Service) NewSegment(segment models.Segment) (error){      
 
}

func (s *Service) DeleteSegment(segment models.Segment) (error){    
 
}