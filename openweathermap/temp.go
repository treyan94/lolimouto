package openweathermap

type Main struct {
	Temp      Temp    `json:"temp"`
	FeelsLike Temp    `json:"feels_like"`
	TempMin   Temp    `json:"temp_min"`
	TempMax   Temp    `json:"temp_max"`
	Pressure  float32 `json:"pressure"`
	Humidity  float32 `json:"humidity"`
}

type Temp float32

func (t Temp) Celsius() float32 {
	return float32(t - 273.15)
}
