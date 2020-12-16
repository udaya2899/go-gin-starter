package configuration

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Configuration exported
type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// ServerConfiguration exported
type ServerConfiguration struct {
	Port int
}

// DatabaseConfiguration exported
type DatabaseConfiguration struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

// New initializes a new Configuration from the ENV variables
func New() *Configuration {
	viper.SetConfigName("config")

	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	var config Configuration

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error while reading config file: %v", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Configuration file changed")
	})

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config file to struct, err: %v", err)
	}

	return &config
}
