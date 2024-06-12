package weather

import (
	"fmt"

	"github.com/vastzp/weather/internal/openweather"
)

type WeatherService struct {
	weatherClient WeatherClient
}

type WeatherClient interface {
	GetWeather(lat, lon float64) (*openweather.WeatherResponse, error)
}

const (
	WeatherThresholdCold     = 32
	WeatherThresholdModerate = 86

	WeatherValueCold     = "cold"
	WeatherValueModerate = "moderate"
	WeatherValueHot      = "hot"
)

type CurrentWeather struct {
	Condition  string // snow, raint, etc
	Feels      string // hot, cold of moderate
	Lat        float64
	Lon        float64
	Celsius    float64
	Fahrenheit float64
	Country    string
}

func NewWeatherService(weatherClient WeatherClient) *WeatherService {
	return &WeatherService{weatherClient: weatherClient}
}

func (ws *WeatherService) GetWeather(lat, lon float64) (*CurrentWeather, error) {
	weatherResp, err := ws.weatherClient.GetWeather(lat, lon)
	if err != nil {
		return nil, fmt.Errorf("failed to interract with weather client: %w", err)
	}

	// Convert response from weather client layer to service layer
	cw := &CurrentWeather{Lat: lat, Lon: lon}
	cw.Condition = weatherResp.Weather[0].Main // todo: be sure that we have at leat one element

	// make a decision how the current temperature feels
	switch {
	case weatherResp.Main.Temp < WeatherThresholdCold:
		cw.Feels = WeatherValueCold
	case weatherResp.Main.Temp < WeatherThresholdModerate:
		cw.Feels = WeatherValueModerate
	default:
		cw.Feels = WeatherValueHot
	}

	// I added this fields just for debug
	cw.Fahrenheit = weatherResp.Main.Temp
	cw.Celsius = (cw.Fahrenheit - 32) * 5 / 9
	cw.Country = weatherResp.Sys.Country

	return cw, nil
}
