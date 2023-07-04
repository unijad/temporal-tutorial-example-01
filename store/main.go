package store

import (
	"context"
	"goenv/messages"
)

func GetCurrentWeather(ctx context.Context, cityName string) (messages.WeatherData, error) {
	// code to fetch the current weather for a given city
	return messages.WeatherData{
		Temperature: 36,
		Humidity:    40,
		WindSpeed:   21,
	}, nil
}
