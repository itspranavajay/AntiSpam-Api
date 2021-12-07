package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment string
	Mongo       config
}

type config struct {
	Server   string
	Database string
}

func (moe *Config) Read() {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&moe)
	if err != nil {
		panic(err)
	}
}
