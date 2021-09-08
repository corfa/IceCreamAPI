package realizationService

import (
	"iceCreamApiWithDI/layers/database"
)

type Service struct {
	database.IDataBase
}

func NewService(db database.IDataBase) *Service {
	return &Service{
		db,
	}
}
