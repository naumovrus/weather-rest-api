package entities

type Weather struct {
	Id     int `json:"id" db:"id"`
	CityId int
	Temp   *float64 `json:"temp" db:"temp"`
}
