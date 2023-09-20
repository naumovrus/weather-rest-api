package entities

type Weather struct {
	Id     int `json:"id" db:"id"`
	CityId int 
	Temp   int `json:"temp" db:"temp"`
}
