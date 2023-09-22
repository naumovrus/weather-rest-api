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
func (w *WeatherAPI) GetWeather(city entities.City) (entities.Weather, error) {
	apiKey := w.ApiKey
	var weather entities.Weather
	request, err := openweathermap.NewCurrent("C", "ru", *apiKey)
	if err != nil {
		return weather, err
	}
	request.CurrentByName(city.Name)
	// cityId := city.Id
	temp := &request.Main.Temp
	weather.Temp = temp
	// weather.CityId = cityId
	return weather, nil
}
