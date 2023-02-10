package main

import (
	"fmt"
	"strings"
        "github.com/Syfaro/telegram-bot-api"
	"github.com/briandowns/openweathermap"
)

type Weather struct {
	Temperature float64 `json:"temp"`
	Humidity    float64 `json:"humidity"`
}

func main() {
	bot, err := tgbotapi.NewBotAPI("YourTelegramToken")
	if err != nil {
		fmt.Println(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !strings.HasPrefix(update.Message.Text, "/weather") {
			continue
		}

		city := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/weather"))
		weather, err := getWeather(city)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			fmt.Sprintf("Темперетура в  %s сейчас %.2f°С, Влажнсть: %.2f%%\n", city, weather.Temperature, weather.Humidity))
		bot.Send(msg)
	}
}

func getWeather(city string) (Weather, error) {
	w, err := openweathermap.NewCurrent("C", "ru", "YourTokenOpenweather")
	if err != nil {
		return Weather{}, err
	}

	w.CurrentByName(city)
	if err != nil {
		return Weather{}, err
	}

	return Weather{
		Temperature: w.Main.Temp,
		Humidity:    float64(w.Main.Humidity),
	}, nil
}
