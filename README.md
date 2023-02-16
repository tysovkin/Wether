<h1 align="center">
# Weather Telegram Bot API
![alt text](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExYzE4M2Y0MzAwMzBmY2IzNDIzNDQ0ZWE1YzMxMDk5OWE4NmJiM2YyMyZjdD1n/qemJG3Zif4NZE6miJ7/giphy.gif)
</h1>

This is a Telegram bot that transmits weather information for the specified city. To get weather information, the OpenWeatherMap API is used

We get the API key (Token) from OpenWeatherMap and Telegram.
Use your own keys (Token) in the program

# How it works

When the user sends a message with the name of the city, the bot checks whether the weather information is in the cache and whether it is valid. If the information is in the cache and it is valid, the bot takes the information. Otherwise, it makes a request to the OpenWeatherMap API to get weather information and stores it in the cache for an hour.

Next, the bot sends the user a message with the weather for the specified city.

## Credits
* [Syfaro/telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api)
* [briandowns/openweathermap](https://github.com/briandowns/openweathermap)

