// Package configuration defines configration structure and methods to handle config
package configuration

import (
	"github.com/spf13/viper"
)

// Configuration defines properties for configration
type Configuration struct {
	// Database config
	Database Database

	// Server config
	Server Server
}

// New creates an instance of Config
func New(configName string, configPath string) (*Configuration, error) {
	config := &Configuration{}
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
