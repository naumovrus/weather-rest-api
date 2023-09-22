package services

import (
	ent "github.com/naumovrus/weather-api/internal/entities"
	"github.com/naumovrus/weather-api/internal/repository"
)

type CitiesSerice struct {
	repo repository.Cities
}

func NewCitiesService(repo repository.Cities) *CitiesSerice {
	return &CitiesSerice{repo: repo}
}

func (s *CitiesSerice) AddCity(userId int, city ent.City) (int, error) {
	return s.repo.AddCity(userId, city)
}

func (s *CitiesSerice) GetAll() ([]ent.City, error) {
	return s.repo.GetAll()
}

func (s *CitiesSerice) GetUsersCity(userId int) ([]ent.City, error) {
	return s.repo.GetUsersCity(userId)
}

func (s *CitiesSerice) GetByName(name string) (ent.City, error) {
	return s.repo.GetByName(name)
}

// func (s *CitiesSerice) Delete()
