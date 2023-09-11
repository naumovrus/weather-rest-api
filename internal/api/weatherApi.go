package api

type WeatherResponse struct {
	Name string  `json:"name"`
	Temp float64 `json:"temp"`
}

func GetWeatherResponse() {

}
