package openweathermap

type City struct {
	Main          `json:"main"`
	WeatherHolder `json:"weather"`
}
