package main

import (
	"log"
	restapi "rest-api"
	"rest-api/pkg/handler"
	"rest-api/pkg/repository"
	"rest-api/pkg/service"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

const (
	url = "http://api.fakeshop-api.com"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "db",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	//зависимости
	repos := repository.NewRepository(db)    //создаем репозиторий
	services := service.NewService(repos)    //создаем сервис который зависит от репозитория
	handlers := handler.NewHandler(services) //создаем обработчик который зависит от сервиса

	srv := new(restapi.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}

// инициализация конфигурационных файлов
func initConfig() error {
	viper.SetConfigName("config") //name of config file
	viper.AddConfigPath("./configs/")
	return viper.ReadInConfig()
}
