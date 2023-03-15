// utils regroups helper methods
package utils

import (
	"os"

	"github.com/spf13/viper"
)

// Config represents the app configuration
type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	SSLMode    string
}

// LoadConfig loads env variables based on env and build a Config object
func LoadConfig(path string, filename string) (config Config, err error) {
	if os.Getenv("ENV") == "production" {
		config := Config{
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			SSLMode:    os.Getenv("SSL_MODE"),
		}
		return config, nil
	} else {
		viper.AddConfigPath(path)
		viper.SetConfigName(filename)
		viper.SetConfigType("env")

		viper.AutomaticEnv()

		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		err = viper.Unmarshal(&config)
		return
	}
}
