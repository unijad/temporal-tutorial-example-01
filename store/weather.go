package store

import (
	"context"
	"goenv/messages"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Weather struct{}

func (w *Weather) GetWeather(ctx context.Context, cityName string) (*messages.WeatherData, error) {
	db := &Database{}
	err := db.Connect()
	if err != nil {
		return nil, err
	}
	data := &messages.WeatherData{
		CityName: cityName,
	}
	record := db.gorm.Where("city_name = ?", cityName).First(data)
	if record.Error != nil {
		logrus.Error(record.Error)
		return nil, record.Error
	}
	return data, nil
}
