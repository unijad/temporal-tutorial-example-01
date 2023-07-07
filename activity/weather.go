package activity

import (
	"context"
	"goenv/messages"
	"goenv/store"
)

func GetWeather(ctx context.Context, cityName string) (*messages.WeatherData, error) {
	// code to fetch the current weather for a given city
	w := &store.Weather{}

	data, err := w.GetWeather(ctx, cityName)
	if err != nil {
		return nil, err
	}

	return data, nil
}
