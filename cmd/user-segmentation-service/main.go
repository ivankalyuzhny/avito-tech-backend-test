package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"

	"avito-tech-backend-test/internal/app/handler"
	"avito-tech-backend-test/internal/app/repository"
	"avito-tech-backend-test/internal/app/service"
)

type Config struct {
	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"database"`
	Server struct {
		Host    string `yaml:"host"`
		Port    int    `yaml:"port"`
		Timeout int    `yaml:"timeout"`
	} `yaml:"server"`
}

func main() {
	config := Config{}
	data, err := os.ReadFile("configs/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlx.Connect(
		config.Database.Driver,
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			config.Database.Host,
			config.Database.Port,
			config.Database.User,
			config.Database.Password,
			config.Database.Name,
			config.Database.Sslmode,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	initSQL, err := os.ReadFile("deployments/init.sql")
	if err != nil {
		log.Panic(err)
	}

	_, err = db.Exec(string(initSQL))
	if err != nil {
		log.Panic(err)
	}

	segmentRepository := repository.NewSegmentRepository(db)
	userRepository := repository.NewUserRepository(db)

	segmentService := service.NewSegmentService(segmentRepository)
	userService := service.NewUserService(userRepository, segmentRepository)

	router := mux.NewRouter()
	handler.RegisterSegmentHandlers(router, segmentService)
	handler.RegisterUserHandlers(router, userService)

	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler:           router,
		ReadHeaderTimeout: time.Duration(config.Server.Timeout) * time.Second,
	}

	log.Printf("Starting server on %s:%d", config.Server.Host, config.Server.Port)
	log.Fatal(server.ListenAndServe())
}
