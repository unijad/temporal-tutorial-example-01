package activity

import (
	"context"
	"goenv/messages"
	"goenv/store"

	"go.temporal.io/sdk/temporal"
)

func GetWeather(ctx context.Context, cityName string) (result messages.WeatherData, err error) {
	result, err = store.GetCurrentWeather(ctx, cityName)
	if err != nil {
		return result, temporal.NewApplicationError("unable to get weather data", "GET_WEATHER", err)
	}
	return result, nil
}
