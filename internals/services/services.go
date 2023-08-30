package services

import (
	// "errors"
 
	// "fmt"
	"avito_test/internals/db"
	"avito_test/internals/models"
	// "github.com/sirupsen/logrus"
)

type Service struct{
	Storage db.Storage
}




func NewService(storage db.Storage) *Service{
	resultService := new(Service)
	resultService.Storage = storage
	return resultService
}



func (s *Service) FindUserSegment(id int64) ([]models.Segment ,error){    
	result,err:=s.Storage.FindUserSegment(id)
	return result,err
}

func (s *Service) NewUserSegment(addRequest models.AddRequest) (error){     
	err := s.Storage.AddUserSegment(addRequest)
	 return err
}
 
func (s *Service) NewSegment(segment models.Segment) ([]models.User, error){      
	result,err := s.Storage.AddSegment(segment)
	return result,err
}

func (s *Service) DeleteSegment(segment models.Segment) (error){    
 return nil
}