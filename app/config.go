package copilot

import (
	"flag"
	"fmt"
	"os"
)

// Config for the API server
type Config struct {
	port      string
	cacheHost string
	cacheURL  string
	// version   bool
	debug bool
	clean bool
}

// Conf global for the API server
var Conf Config

func parseConfig() (Config, error) {
	var conf = Config{}
	var clean, debug bool
	// version,
	// flag.BoolVar(&version, "version", false, "API Version")
	flag.BoolVar(&clean, "clean", false, "Clean cache")
	flag.BoolVar(&debug, "debug", false, "Debug API")
	flag.Parse()

	conf.port = getEnv("PORT", "8080")
	conf.cacheHost = getEnv("CACHE_HOST", "redis")
	conf.cacheURL = fmt.Sprintf("redis://%s:6379/0", conf.cacheHost)
	// conf.version = version
	conf.debug = debug
	conf.clean = clean

	if debug != false {
		// app.StartProfiler()
	}

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
