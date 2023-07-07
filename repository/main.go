package repository

import (
	"goenv/messages"

	_ "github.com/lib/pq"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DB_USER     = "1"
	DB_PASSWORD = ""
	DB_NAME     = "temporal_example"
)

type Repository struct {
	gorm *gorm.DB
}

func (d *Repository) Connect() error {
	db, err := gorm.Open(sqlite.Open("store.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&messages.WeatherData{})
	if err != nil {
		panic("failed to migrate database")
	}
	// look if any records exists and if not add record
	var count int64
	db.Model(&messages.WeatherData{}).Count(&count)
	if count == 0 {
		db.Create(&messages.WeatherData{CityName: "London", Temperature: 36, Humidity: 40, WindSpeed: 21})
	}

	d.gorm = db

	return nil
}
