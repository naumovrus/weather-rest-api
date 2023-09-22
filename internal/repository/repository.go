package repository

import (
	"github.com/jmoiron/sqlx"

	ent "github.com/naumovrus/weather-api/internal/entities"
)

type Authorization interface {
	CreateUser(user ent.User) (int, error)
	GetUser(username, password string) (ent.User, error)
}

type Cities interface {
	GetByName(name string) (ent.City, error)
	AddCity(userId int, city ent.City) (int, error)
	GetUsersCity(userId int) ([]ent.City, error)
	GetAll() ([]ent.City, error)
	DeleteCity(userId int, cityId int) error
}

type Weather interface {
	AddWeather(cityId int, weather ent.Weather) (int, error)
	GetWeather(cityId int) (ent.Weather, error)
}

type Repository struct {
	Authorization
	Cities
	Weather
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Cities:        NewCityPostrgres(db),
		Weather:       NewWeatherPostgres(db),
	}

}
