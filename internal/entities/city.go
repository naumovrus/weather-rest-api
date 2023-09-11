package entities

type City struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" binding:"required"`
}
