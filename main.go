package main

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/briandowns/openweathermap"
	"strings"
)

// Weather структура для хранения информации о температуре и влажности
type Weather struct {
	Temperature float64 `json:"temp"`
	Humidity    float64 `json:"humidity"`
}

func main() {
	// Создаем Telegram bot API
	bot, err := tgbotapi.NewBotAPI("YourTelegramToken")
	if err != nil {
		fmt.Println(err)
	}

	// включакм режим отладки
	bot.Debug = true

	//Настроим обновления с таймаутом в 60 секунд
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	// Получаем обновление от бота
	updates, err := bot.GetUpdatesChan(updateConfig)

	//Циклический просмотр обновлений
	for update := range updates {
		//Пропускаем обновления без сообщений
		if update.Message == nil {
			continue
		}

		// Проверяем, начинается ли сообщение с "/weather".
		if !strings.HasPrefix(update.Message.Text, "/weather") {
			continue
		}

		// Получаем название города из сообщения
		city := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, "/weather"))
		//Получаем информацию о погоде в городе
		weather, err := getWeather(city)
		if err != nil {
			// Отправляет сообщение об ошибке, если возникла проблема с получением информации о погоде
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
			continue
		}

		// Создаем возвращаемое сообщение с информацией о погоде
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			fmt.Sprintf("Темперетура в  %s сейчас %.2f°С, Влажнсть: %.2f%%\n", city, weather.Temperature, weather.Humidity))
		// Отпровляем сообщение 
		bot.Send(msg)
	}
}

// getWeather возвращает текущую информацию о погоде для города
func getWeather(city string) (Weather, error) {
	// Create a new instance of the OpenWeatherMap API
	w, err := openweathermap.NewCurrent("C", "ru", "YourTokenOpenweather")
	if err != nil {
		return Weather{}, err
	}

	//Получаем текущую информацию о погоде в городе
	w.CurrentByName(city)
	if err != nil {
		return Weather{}, err
	}

	// Возвращаем информацию о температуре и влажности в виде структуры 
	return Weather{
		Temperature: w.Main.Temp,
		Humidity:    float64(w.Main.Humidity),
	}, nil
}

