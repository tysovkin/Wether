package main

import (
	"encoding/json"
	"fmt"
	tele "github.com/tucnak/telebot"
	"io"
	"net/http"
	"strings"
	"time"
)

type Weather struct {
	Temp     float64 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

func main() {
	b, err := tele.NewBot(tele.Settings{
		Token:  "{YourTelegramToken}",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	b.Handle("/weather", func(message tele.Message) {
		city := strings.TrimSpace(message.Text[len("/weather"):])

		if city == "" {
			b.Send(message.OriginalChat, "Укажите город...", nil)
			return
		}

		weather, err := getWeather(city)
		if err != nil {
			b.Send(message.Chat, "Нет Данных о Погоде"+err.Error(), nil)
			return
		}

		b.Send(message.Chat,
			fmt.Sprintf("Temperature in %s: %.1f°C\nPressure: %d hPa\nHumidity: %d%%\nMin: %.1f°C\nMax: %.1f°C",
				city, weather.Temp-273.15, weather.Pressure, weather.Humidity, weather.TempMin-273.15, weather.TempMax-273.15), nil)
	})

	b.Start()
	if err != nil {
		fmt.Println(err)
	}

}

func getWeather(city string) (*Weather, error) {
	resp, err := http.Get("http://api.openweathermap.org/geo/1.0/weather?q=" + city + "&limit=5&appid={YourOpenWetherToken}\n")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather Weather
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, err
	}

	return &weather, nil
}
