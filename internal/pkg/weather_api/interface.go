package weatherapi

import "github.com/naumovrus/weather-api/internal/entities"

type WeatherApi interface {
	GetWeather(city *string) (*entities.Weather, error) 
}
