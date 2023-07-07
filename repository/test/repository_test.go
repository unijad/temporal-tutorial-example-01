package test

import (
	"context"
	"goenv/repository"
	"testing"
)

func TestRepository(t *testing.T) {
	w := &repository.Weather{}
	record, err := w.GetWeather(context.Background(), "London")
	if err != nil {
		t.Fatal(err)
	}
	if record.CityName != "London" {
		t.Fatalf("expected London but got %s", record.CityName)
	}
	if record.Temperature != 36 {
		t.Fatalf("expected 36 but got %f", record.Temperature)
	}
	if record.Humidity != 40 {
		t.Fatalf("expected 40 but got %f", record.Humidity)
	}
	if record.WindSpeed != 21 {
		t.Fatalf("expected 21 but got %f", record.WindSpeed)
	}
}
