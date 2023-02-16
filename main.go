package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/briandowns/openweathermap"
	"strings"
	"time"
)

type Weather struct {
	Temperature float64 `json:"temp"`
	Humidity    float64 `json:"humidity"`
	Clouds      float64 `json:"clouds"`
	Rain        float64 `json:"rain"`
}

var cache = make(map[string]Weather)
var cacheLife = make(map[string]time.Time)

func main() {
	bot, err := tgbotapi.NewBotAPI("YourTelegramToken")
	if err != nil {
		fmt.Println(err)
		return
	}
	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	updates, err := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	format := "температура %s: %.1f°C; облачность %.f%%; влажность %.1f%%; дождь в час %.fmm"

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() && update.Message.Command() == "start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Укажите город")
			bot.Send(msg)
			continue
		}

		city := strings.TrimSpace(update.Message.Text)
		weather, err := getWeather(city)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(format,
			city, weather.Temperature, weather.Clouds, weather.Humidity, weather.Rain))
		bot.Send(msg)
	}
}

func getWeather(city string) (Weather, error) {
	if weather, ok := cache[city]; ok && time.Now().Before(cacheLife[city]) {
		return weather, nil
	}

	w, err := openweathermap.NewCurrent("C", "ru", "YourOpenweatherToken")
	if err != nil {
		return Weather{}, err
	}
	w.CurrentByName(city)
	if err != nil {
		return Weather{}, err
	}

	weather := Weather{
		Temperature: w.Main.Temp,
		Humidity:    float64(w.Main.Humidity),
		Clouds:      float64(w.Clouds.All),
		Rain:        w.Rain.OneH,
	}

	cache[city] = weather
	cacheLife[city] = time.Now().Add(time.Hour)

	return weather, nil
}


