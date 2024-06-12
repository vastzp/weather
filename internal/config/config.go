package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config represents a config.
type Config struct {
	APIConfig *APIConfig
	OW        *OpenWeatherConfig
}

// OpenWeatherConfig - config for the OpenWeather client
type OpenWeatherConfig struct {
	APIKey string
}

// APIConfig - web server configuration
type APIConfig struct {
	Host string
	Port int
}

func New() (*Config, error) {
	c := &Config{}

	// API configurations
	host, err := getEnvStr("SERVER_HOST")
	if err != nil {
		return nil, err
	}

	port, err := getEnvInt("SERVER_PORT")
	if err != nil {
		return nil, err
	}

	c.APIConfig = &APIConfig{
		Host: host,
		Port: port,
	}

	// OpenWeather configuration
	owAPIKEY, err := getEnvStr("OPEN_WEATHER_API_KEY")
	if err != nil {
		return nil, err
	}
	c.OW = &OpenWeatherConfig{APIKey: owAPIKEY}

	return c, nil

}

func getEnvStr(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if exists {
		value = strings.TrimSpace(value)
		return value, nil
	}
	return "", fmt.Errorf("environment variable '%s' isn't set", key)
}

func getEnvInt(key string) (int, error) {
	val, err := getEnvStr(key)
	if err != nil {
		return 0, err
	}

	v, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("environment variable '%s' must contains an integer value", key)
	}

	return v, nil
}
