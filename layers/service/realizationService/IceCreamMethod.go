package realizationService

import (
	"errors"
	"iceCreamApiWithDI/layers/handler/ModelForGin"
)

func (s *Service) AppendIceCream(data ModelForGin.IceCream, token ModelForGin.HeaderAuth)error  {
	_ , errToken:=s.ReadJWToken(token.Authorization)
	if errToken != nil{
		return errors.New("wrong token")
	}
	errDb := s.IDataBase.CreateIceCream(data)
	return errDb

}
func (s *Service) DeleteIceCream(data ModelForGin.IceCream, token ModelForGin.HeaderAuth) error  {
	_ , errToken:=s.ReadJWToken(token.Authorization)
	if errToken != nil{
		return errors.New("wrong token")
	}
	errDb := s.IDataBase.DeleteIceCream(data)
	return errDb
}
