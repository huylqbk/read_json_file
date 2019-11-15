package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const ConfigFile = "./config/"

type Config struct {
	Data Data `toml:"data"`
}

type Data struct {
	Organization string `toml:"Organization"`
	User         string `toml:"User"`
	Ticket       string `toml:"Ticket"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("default")
	viper.SetConfigType("toml")
	viper.AddConfigPath(ConfigFile)
	viper.ReadInConfig()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	var conf *Config
	err := viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
