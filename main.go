package main

import (
	"iceCreamApiWithDI/database"
	"iceCreamApiWithDI/database/ModelForGorm"
	"iceCreamApiWithDI/handler"
	"iceCreamApiWithDI/server"
	"iceCreamApiWithDI/service"
	"log"
)

func main() {

	configDB := database.ConfigDB{
		Password: "00000",
		Host:     "localhost",
		Sslmode:  "disable",
		Port:     "4000",
		User:     "postgres",
	}
	EngineDatabase, errDBacon := database.EnginPostgres(configDB)
	if errDBacon != nil {
		log.Panicln(errDBacon)
	}

	MigrateFlag := false

	if MigrateFlag != false {
		EngineDatabase.AutoMigrate(&ModelForGorm.Users{}, &ModelForGorm.IceCreams{})
	}

	Repository := database.NewDataBase(EngineDatabase)
	Service := service.NewServices(Repository)
	MyHandler := handler.NewHandler(Service)
	Server := new(server.Server)
	errSrv := Server.Run("8080", MyHandler.InitHandler())
	if errSrv != nil {
		log.Panicln("srv Error")
	}
}
