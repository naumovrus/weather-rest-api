package repository

import (
	"github.com/jmoiron/sqlx"
	ent "github.com/naumovrus/weather-api/internal/entities"
)

const (
	weatherTable = "weather"
)

type WeatherPostgres struct {
	db *sqlx.DB
}

func NewWeatherPostgres(db *sqlx.DB) *WeatherPostgres {
	return &WeatherPostgres{db: db}
}

func (r *WeatherPostgres) AddWeather(cityId int, weather ent.Weather) (int, error) {
	// query := fmt.Sprintf("INSERT INTO %s (city_id, temp) VALUES ($1, $2) RETURNING id", weatherTable)
	return 0, nil
}

func (r *WeatherPostgres) GetWeather(cityId int) (ent.Weather, error) {
	return ent.Weather{}, nil
}
