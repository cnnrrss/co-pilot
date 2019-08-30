package copilot

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Config for the API server
type Config struct {
	port      string
	cacheHost string
	cacheURL  string
	apiKey    string
	clean     bool
}

// Conf global for the API server
var Conf Config

func parseConfig() (Config, error) {
	var conf = Config{}

	conf.apiKey = getEnv("API_KEY", "")
	if conf.apiKey == "" {
		return conf, errors.New("api key cannot be empty")
	}

	conf.port = getEnv("PORT", "8080")
	conf.cacheHost = getEnv("CACHE_HOST", "redis")
	conf.cacheURL = fmt.Sprintf("redis://%s:6379/0", conf.cacheHost)
	clean, _ := strconv.ParseBool(getEnv("CACHE_FLUSH", "FALSE"))
	conf.clean = clean

	return conf, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// SetDefaultConfig parses and assigns the config
func SetDefaultConfig() error {
	conf, err := parseConfig()
	if err != nil {
		return err
	}
	Conf = conf
	return nil
}
