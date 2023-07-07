package activity

import (
	"context"
	"goenv/messages"
	"goenv/repository"
)

func GetWeather(ctx context.Context, cityName string) (*messages.WeatherData, error) {
	// code to fetch the current weather for a given city
	w := &repository.Weather{}

	data, err := w.GetWeather(ctx, cityName)
	if err != nil {
		return nil, err
	}

	return data, nil
}
