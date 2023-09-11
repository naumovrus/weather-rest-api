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
	AddCity(userId int, city ent.City) (int, error)
	GetUsersCity(userId int) ([]ent.City, error)
	GetAll() ([]ent.City, error)
	// Delete(userId int, cityId int) error
}

type Repository struct {
	Authorization
	Cities
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Cities:        NewCityPostrgres(db),
	}

}
