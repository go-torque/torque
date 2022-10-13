package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var conf *viper.Viper

type Config interface {
}

func Load() {
	conf = viper.New()

	conf.SetConfigFile(".env")
	conf.SetConfigType("env")

	err := conf.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
