package main

import (
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/naumovrus/weather-api/internal/config"
	handler "github.com/naumovrus/weather-api/internal/handlers"
	httpserver "github.com/naumovrus/weather-api/internal/http-server"
	weatherapi "github.com/naumovrus/weather-api/internal/pkg/weather_api"
	"github.com/naumovrus/weather-api/internal/repository"
	"github.com/naumovrus/weather-api/internal/services"
	"github.com/sirupsen/logrus"

	"golang.org/x/exp/slog"
)

var (
	ApiKey = os.Getenv("OPEN_WEATHER_MAP_API")
)

func main() {
	// init config
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error load .env files: %s", err.Error())
	}
	cfg := config.LoadConfig()
	// logger := initLogger()
	apiKey := os.Getenv("OMW_API")
	//init db

	db, err := repository.NewPostgresDB(repository.Config{
		Username: cfg.Username,
		Host:     cfg.Host,
		Port:     cfg.PortDb,
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   cfg.Dbname,
		SSLMode:  cfg.Sslmode,
	})
	if err != nil {
		logrus.Fatalf("unabled to initialize db: %s", err.Error())
	}

	weatapi := weatherapi.NewWeatherApi(&apiKey)
	repos := repository.NewRepository(db)
	services := services.NewService(repos)
	handlers := handler.NewHandler(services, weatapi)
	srv := new(httpserver.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initLogger() *slog.Logger {

	log := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	return log
}
