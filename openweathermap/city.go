package openweathermap

import "fmt"

type City struct {
	Main          `json:"main"`
	WeatherHolder `json:"weather"`
}

func (c City) Text() string {
	if len(c.WeatherHolder) == 0 {
		return "Error, please try again later."
	}

	w := c.WeatherHolder[0]
	str := fmt.Sprintf("%s %.2f°C %s", w.Emoji(), c.Main.Temp.Celsius(), w.Description)
	return str
}
