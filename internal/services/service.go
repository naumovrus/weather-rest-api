package services

import (
	ent "github.com/naumovrus/weather-api/internal/entities"
	"github.com/naumovrus/weather-api/internal/repository"
)

type Authorization interface {
	CreateUser(user ent.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Cities interface {
	AddCity(userId int, city ent.City) (int, error)
	GetUsersCity(userId int) ([]ent.City, error)
	GetAll() ([]ent.City, error)
	// Delete(userId int, cityId int) error
}

type Service struct {
	Authorization
	Cities
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Cities:        NewCitiesService(repos.Cities),
	}
}
