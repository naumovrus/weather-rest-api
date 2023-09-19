package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	ent "github.com/naumovrus/weather-api/internal/entities"
)

const (
	cityiesTable   = "cities"
	usersCityTable = "users_city"
)

type CityPostgres struct {
	db *sqlx.DB
}

func NewCityPostrgres(db *sqlx.DB) *CityPostgres {
	return &CityPostgres{db: db}
}

func (r *CityPostgres) AddCity(userId int, city ent.City) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int

	name := city.Name
	checkCityQuerry := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", cityiesTable)
	row := r.db.QueryRow(checkCityQuerry, name)
	if err := row.Scan(&id); err != nil {
		// Case when we do not have input city in table cityies
		addCityQuery := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", cityiesTable)
		row = tx.QueryRow(addCityQuery, name)
		if err := row.Scan(&id); err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	// Case when we had or when we created input city in table cityies
	createUsersCityQuery := fmt.Sprintf("INSERT INTO %s (user_id, city_id) VALUES ($1, $2)", usersCityTable)
	_, err = tx.Exec(createUsersCityQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *CityPostgres) GetByName(name string) (ent.City, error) {
	var city ent.City
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1", usersTable)
	err := r.db.Get(&city, query, name)
	return city, err
}

func (r *CityPostgres) GetAll() ([]ent.City, error) {
	var cityies []ent.City
	query := fmt.Sprintf("SELECT name FROM %s", cityiesTable)
	err := r.db.Select(&cityies, query)
	return cityies, err
}

func (r *CityPostgres) GetUsersCity(userId int) ([]ent.City, error) {
	var cityies []ent.City
	query := fmt.Sprintf(`SELECT name FROM %s cl INNER JOIN %s 
	uc on cl.id = uc.city_id WHERE uc.user_id = $1`,
		cityiesTable, usersCityTable)
	err := r.db.Select(&cityies, query, userId)
	return cityies, err
}

func (r *CityPostgres) DeleteCity(userId int, cityId int) error {
	query := fmt.Sprintf(`DELETE FROM %s uc WHERE user_id = $1 AND city_id = $2`, usersCityTable)
	_, err := r.db.Exec(query, userId, cityId)
	return err
}
