package realizationService

import (
	"errors"
	ModelForGin2 "iceCreamApiWithDI/layers/handler/ModelForGin"
)

func (s *Service) CreateUser(data ModelForGin2.CreateUser) error {
	data.Password = s.HashPassword(data.Password)
	err := s.IDataBase.CreateUser(data)
	return err
}

func (s *Service) LoginUser(data ModelForGin2.GetUser) (string, error) {
	user, userErr := s.IDataBase.GetUser(data)
	if userErr != nil {
		return "", userErr
	}
	flag := s.CheckHashPassword(data.Password, user.Password)
	if flag != true {
		return "", errors.New("wrong password")
	}
	return s.CreateJWToken(user.Id)

}


