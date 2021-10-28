package openweathermap

type WeatherHolder []Weather

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

func (w Weather) Emoji() string {
	switch w.Icon {
	case "01d":
		return "☀️"
	case "01n":
		return "🌃"
	case "02d":
		return "🌤️"
	case "02n":
		return "🌁"
	case "03d":
		fallthrough
	case "03n":
		fallthrough
	case "04d":
		fallthrough
	case "04n":
		return "☁️"
	case "09d":
		fallthrough
	case "09n":
		fallthrough
	case "10n":
		return "🌧️"
	case "10d":
		return "🌦️"
	case "11d":
		fallthrough
	case "11n":
		return "⛈️"
	case "13d":
		fallthrough
	case "13n":
		return "🌨️"
	case "50d":
		fallthrough
	case "50n":
		return "🌫️"
	default:
		return "❓"
	}
}
