package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/tucnak/telebot"
)

type Weather struct {
	Temp     float64 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_min"`
	TempMax  float64 `json:"temp_max"`
}

func main() {
	b, err := telebot.NewBot(telebot.Settings{
		Token:  "Mytoken",
		Poller: &telebot.LongPoller{10 * time.Second},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	b.Handle("/weather", func(message *telebot.Message) {
		city := strings.TrimSpace(message.Text[len("/weather"):])

		if city == "" {
			b.Send(message.Chat, "Укажите город...", nil)
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
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=OpenWeatherAPI\n")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weather Weather
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

