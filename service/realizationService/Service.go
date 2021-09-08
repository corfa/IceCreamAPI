package realizationService

import "iceCreamApiWithDI/database"

type Service struct {
	database.IDataBase
}

func NewService(db database.IDataBase) *Service {
	return &Service{
		db,
	}
}
