package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// env variables struct
type Config struct {
	PORT         string
	POSTGRES_URI string
}

func SetConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("dotenv")

	// read the config file
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	port := viper.GetString("PORT")
	postgresURI := viper.GetString("POSTGRES_URI")

	if port == "" || postgresURI == "" {
		return nil, fmt.Errorf("missing required env variables")
	}
	return &Config{
		PORT:         port,
		POSTGRES_URI: postgresURI,
	}, nil
}
