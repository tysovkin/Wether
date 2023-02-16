# Weather Telegram Bot API

![alt text](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExYzE4M2Y0MzAwMzBmY2IzNDIzNDQ0ZWE1YzMxMDk5OWE4NmJiM2YyMyZjdD1n/qemJG3Zif4NZE6miJ7/giphy.gif)


Это Telegram-бот, который передает информацию о погоде для указаного города. Для получения информации о погоде используется OpenWeatherMap API

Получаем API-ключ(Token) от OpenWeatherMap и Telegram.
В програме импользуйте свои ключи(Token) 

Как это работает

Когда пользователь отправляет сообщение с названием города, бот проверяет наличие информация о погоде в кэше и действительна ли она. Если информация в кэше и она действительна, бот забирает информацию. В противном случае делает запрос к API OpenWeatherMap для получения информации о погоде и сохраняет на час в кэше.

Далее бот отправляет пользователю сообщение с погодой для указанного города.

Зависимости 
github.com/Syfaro/telegram-bot-api - Telegram Bot API
github.com/briandowns/openweathermap - OpenWeatherMap API wrapper

