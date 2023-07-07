package messages

// define the workflow function
type WeatherData struct {
	CityName    string
	Temperature float64
	Humidity    float64
	WindSpeed   float64
}

type WeatherSignal struct {
	Message string
}

type Error struct {
	RunId   string
	Message string
}

func (e *Error) ErrorJSON() []byte {
	return []byte(`{"run_id": "` + e.RunId + `", "message": "` + e.Message + `"}`)
}
