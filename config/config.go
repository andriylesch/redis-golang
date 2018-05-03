package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	// AppPort ...
	AppPort string
	// SvcTracingZipkin ...
	SvcTracingZipkin string
	// RedisHost ...
	RedisHost string
)

func init() {
	// Get the config
	InitConfig()
}

// New gets the service configuration
func InitConfig() {
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("REDIS_HOST", "localhost:6379")
	viper.SetDefault("SVC_TRACING_ZIPKIN", "http://localhost:9411")

	if os.Getenv("ENVIRONMENT") == "DEV" {
		_, dirname, _, _ := runtime.Caller(0)
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(filepath.Dir(dirname))
		viper.ReadInConfig()
	} else {
		viper.AutomaticEnv()
	}

	// Assign env variables value to global variables
	AppPort = viper.GetString("APP_PORT")
	RedisHost = viper.GetString("REDIS_HOST")
	SvcTracingZipkin = viper.GetString("SVC_TRACING_ZIPKIN")

}
