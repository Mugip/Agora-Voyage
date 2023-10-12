package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// Config represents the application configuration.
type Config struct {
	Environment string
	// Add other required configuration fields here
}

// Load loads the configuration from environment variables or configuration files
func (c *Config) Load() error {
	if err := c.loadFromEnv(); err != nil {
		return err
	}

	if err := c.loadFromFile(); err != nil {
		return err
	}

	// Add any additional validation or processing of configuration here

	return nil
}

// loadFromEnv loads the configuration from environment variables
func (c *Config) loadFromEnv() error {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	c.Environment = viper.GetString("environment")
	// Load other environment variables here

	return nil
}

// loadFromFile loads the configuration from a configuration file, such as YAML, JSON, or TOML
func (c *Config) loadFromFile() error {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		exePath, err := os.Executable()
		if err != nil {
			return err
		}
		configPath = filepath.Dir(exePath)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	return viper.ReadInConfig()
}
