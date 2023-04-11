package main

import (
	restapi "rest-api"
	"rest-api/pkg/handler"
	"rest-api/pkg/repository"
	"rest-api/pkg/service"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

/*const (
	url = "http://api.fakeshop-api.com"
)*/

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		//Password: os.Getenv("DB_PASSWORD"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	//зависимости
	repos := repository.NewRepository(db)    //создаем репозиторий
	services := service.NewService(repos)    //создаем сервис который зависит от репозитория
	handlers := handler.NewHandler(services) //создаем обработчик который зависит от сервиса

	srv := new(restapi.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

// инициализация конфигурационных файлов
func initConfig() error {
	viper.SetConfigName("config") //name of config file
	viper.AddConfigPath("./configs/")
	return viper.ReadInConfig()
}
