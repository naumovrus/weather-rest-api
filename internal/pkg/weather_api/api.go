package weatherapi

import (
	"github.com/briandowns/openweathermap"
	"github.com/naumovrus/weather-api/internal/entities"
)

type WeatherAPI struct {
	ApiKey *string
}

func NewWeatherApi(apiKey *string) *WeatherAPI {
	return &WeatherAPI{
		ApiKey: apiKey,
	}
}

// Посмотреть по аналогии с имплиметацией методов сервиса
func (w *WeatherAPI) GetWeather(city entities.City) (*float64, error) {
	apiKey := w.ApiKey
	request, err := openweathermap.NewCurrent("C", "ru", *apiKey)
	if err != nil {
		return nil, err
	}

	request.CurrentByName(city.Name)
	// var weather entities.Weather
	// cityId := city.Id
	temp := &request.Main.Temp
	// weather.CityId = cityId
	// weather.Temp = *temp
	return temp, nil
}
