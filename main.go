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
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
}

func main() {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  "MyTokenBot",
		Poller: &telebot.LongPoller{10 * time.Second},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.Handle("/weather", func(message *telebot.Message) {
		city := strings.TrimSpace(message.Text[len("/weather"):])

		if city == "" {
			bot.Send(message.Chat, "Укажите город...", nil)
			return
		}

		weather, err := getWeather(city)
		if err != nil {
			bot.Send(message.Chat, "Нет Данных о Погоде"+err.Error(), nil)
			return
		}

		bot.Send(message.Chat,
			fmt.Sprintf("Temperature in %s: %.1f°C\nPressure: %d hPa\nHumidity: %d%%\nMin: %.1f°C\nMax: %.1f°C",
				city, weather.Main.Temp-273.15, weather.Main.Pressure, weather.Main.Humidity, weather.Main.TempMin-273.15, weather.Main.TempMax-273.15), nil)
	})

	err = bot.Start()
	if err != nil {
		fmt.Println(err)
	}
}

func getWeather(city string) (*Weather, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=APIopeneWather")
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
