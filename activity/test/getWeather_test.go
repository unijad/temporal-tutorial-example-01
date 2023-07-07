package test

import (
	"context"
	"goenv/activity"
	"testing"
)

func TestGetWeather(t *testing.T) {
	ctx := context.Background()
	cityName := "London"
	weather, err := activity.GetWeather(ctx, cityName)
	if err != nil {
		t.Fatal(err)
	}
	if weather.CityName != cityName {
		t.Fatalf("expected %s but got %s", cityName, weather.CityName)
	}
	if weather.Temperature != 36 {
		t.Fatalf("expected 36 but got %f", weather.Temperature)
	}
	if weather.Humidity != 40 {
		t.Fatalf("expected 40 but got %f", weather.Humidity)
	}
	if weather.WindSpeed != 21 {
		t.Fatalf("expected 21 but got %f", weather.WindSpeed)
	}
}
