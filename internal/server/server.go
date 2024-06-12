package server

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/vastzp/weather/internal/service/weather"
)

type Server struct {
	weatherService weather.WeatherService
	log            *slog.Logger
}

type WrongRequest struct {
	Message string
}

func (s *Server) WeatherHandler(c echo.Context) error {

	sLat := c.QueryParam("lat")
	sLon := c.QueryParam("lon")

	s.log.Debug("new weather request", slog.String("latitude", sLat), slog.String("longtitude", sLon))

	lat, err := strconv.ParseFloat(sLat, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &WrongRequest{Message: "wrong latitude value"})
	}

	lon, err := strconv.ParseFloat(sLon, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &WrongRequest{Message: "wrong longitude value"})
	}

	currentWeather, err := s.weatherService.GetWeather(lat, lon)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &WrongRequest{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, currentWeather)
}

func NewServer(weatherService weather.WeatherService, log *slog.Logger) *Server {
	return &Server{weatherService: weatherService, log: log}
}
