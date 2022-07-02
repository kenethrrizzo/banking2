package config

import "github.com/spf13/viper"

var vp *viper.Viper

func LoadConfig() (*Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName("config")
	vp.AddConfigPath(".")
	vp.SetConfigType("json")

	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := vp.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

type DBConfig struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type Config struct {
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
}
