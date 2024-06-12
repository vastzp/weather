package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/labstack/echo"
	"github.com/vastzp/weather/internal/config"
	"github.com/vastzp/weather/internal/openweather"
	"github.com/vastzp/weather/internal/server"
	"github.com/vastzp/weather/internal/service/weather"
)

func main() {

	// Load config
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	// Logger
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	log.Info("starting weather service", slog.String("host", cfg.APIConfig.Host), slog.Int("port", cfg.APIConfig.Port))

	owClient := openweather.NewOpenWeatherClient(cfg.OW.APIKey)
	weatherService := weather.NewWeatherService(owClient)

	// Web servier. I used echo because it was mentioned in the conversation with Prateep.
	e := echo.New()
	e.Debug = true
	srv := server.NewServer(*weatherService, log)

	e.GET("/weather", srv.WeatherHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", cfg.APIConfig.Host, cfg.APIConfig.Port)))
}
