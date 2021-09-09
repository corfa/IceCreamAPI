package main

import (
	"github.com/joho/godotenv"
	"iceCreamApiWithDI/layers/database"
	"iceCreamApiWithDI/layers/database/ModelForGorm"
	"iceCreamApiWithDI/layers/handler"
	"iceCreamApiWithDI/layers/service"
	"iceCreamApiWithDI/server"
	"log"
	"os"
	"strconv"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	configDB := database.ConfigDB{
		Password: os.Getenv("DB_password"),
		Host:     os.Getenv("DB_host"),
		Sslmode:  os.Getenv("DB_sslmode"),
		Port:     os.Getenv("DB_port"),
		User:     os.Getenv("DB_user"),
	}
	EngineDatabase, errDBacon := database.EnginPostgres(configDB)
	if errDBacon != nil {
		log.Panicln(errDBacon)
	}

	needToMig,_:=strconv.ParseBool(os.Getenv("DB_need_to_migrate"))

	if needToMig {
		errMig:=EngineDatabase.AutoMigrate(&ModelForGorm.Users{}, &ModelForGorm.IceCreams{})
		if errMig!=nil{
			log.Panicln(errMig)
		}
	}

	Repository := database.NewDataBase(EngineDatabase)
	Service := service.NewServices(Repository)
	MyHandler := handler.NewHandler(Service)
	Server := new(server.Server)
	errSrv := Server.Run(os.Getenv("ServerPort"), MyHandler.InitHandler())
	if errSrv != nil {
		log.Panicln("server Error")
	}
}
