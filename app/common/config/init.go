package config

import (
	"github.com/spf13/viper"
)

type ConfigType struct {
}

var Config ConfigType

func Init() {
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic("Error loading .env file")
	}

	if err := viper.Unmarshal(&Config); err != nil {
		panic("Error unmarshal .env file")
	}
}
