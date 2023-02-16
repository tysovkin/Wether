# Weather Telegram Bot API

![alt text](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExYzE4M2Y0MzAwMzBmY2IzNDIzNDQ0ZWE1YzMxMDk5OWE4NmJiM2YyMyZjdD1n/qemJG3Zif4NZE6miJ7/giphy.gif)


package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/briandowns/openweathermap"
	"strings"
	"time"
)

  создаем Структуру Weather c информацией о погоде для города
type Weather struct {
	Temperature float64 `json:"temp"`
	Humidity    float64 `json:"humidity"`
	Clouds      float64 `json:"clouds"`
	Rain        float64 `json:"rain"`
}

// кэш - это map, на которой хранится информация о погоде (the cache is a map on which weather information is stored)
var cache = make(map[string]Weather)
var cacheLife = make(map[string]time.Time)

func main() {
	//  устанавливаем для бота API key (install the Api key for bot)
	bot, err := tgbotapi.NewBotAPI("YourTelegramToken")
	if err != nil {
		fmt.Println(err)
	}
	// включить режим отладки бота ( enable debug mode for bot)
	bot.Debug = true
	// создаем обновление конфигурации ( creating a configuration update)
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60
	// получаем канал обновлений для бота ( get an update channel for the bot )
	updates, err := bot.GetUpdatesChan(updateConfig)
	// для update получаем каждый элемент из updates
	for update := range updates {
		//пропускаем обновление если сообшение nil (skip the update if the message is nil)
		if update.Message == nil {
			continue
		}

		// обрезаем лишнее в название города(cut off the excess in the name of the city)
		city := strings.TrimSpace(update.Message.Text)
		//узнаем погоду для города(find out the weather for the city)
		weather, err := getWeather(city)
		if err != nil {
			//если произошла ошибка, отправляем сообщение об ошибке пользователю(if an error has occurred, we send an error message to the user)
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
			continue
		}
		
		// создаем сообщение с информацией(creating a message with information)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			fmt.Sprintf("Температура  %s  %.1f°С;  Облачность %.f%%;   Влажность %.1f%%;    Дождь за час  %.fmm  \n",
			city, weather.Temperature, weather.Clouds, weather.Humidity, weather.Rain, ))
		
		//отправляем сообщение пользователю
		bot.Send(msg)
	}
}

// функция извлекает погоду для города(function extracts the weather for the city)
func getWeather(city string) (Weather, error) {
	// Проверяем есть ли погода для города в кэше и действительна ли она по-прежнему(checking weather for the city is in the cache and whether it is valid)
	if weather, ok := cache[city]; ok {
		if time.Now().Before(cacheLife[city]) {
			return weather, nil
		}
	}
	//если погода отсутствует в кэше или срок ее истек, получить погоду из API(if the weather is not in the cache or it has expired, get the weather from the AP)
	w, err := openweathermap.NewCurrent("C", "ru", "YourOpenweatherToken")
	if err != nil {
		return Weather{}, err
	}
	// получаем погоду для города(getting the weather for the city)
	w.CurrentByName(city)
	if err != nil {
		return Weather{}, err
	}
	//создаем структуру Weather с полученной информацией погоды(creating a Weather structure with the received weather information)
	weather := Weather{
		Temperature: w.Main.Temp,
		Humidity:    float64(w.Main.Humidity),
		Clouds:      float64(w.Clouds.All),
		Rain:        float64(w.Rain.OneH),
	}

	// Добавим weather в cashe и установим время хранения(Add weather to the cache and set the storage time)
	cache[city] = weather
	cacheLife[city] = time.Now().Add(time.Hour * 1)

	return weather, nil
}

