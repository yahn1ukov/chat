package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP struct {
		Port int `yaml:"port"`
	} `yaml:"http"`

	DB struct {
		Postgres struct {
			URL string `yaml:"url"`
		} `yaml:"postgres"`
	} `yaml:"db"`

	NATS struct {
		Addr string `yaml:"addr"`
	} `yaml:"nats"`

	Session struct {
		Secret string `yaml:"secret"`
	} `yaml:"session"`
}

func New(path string) (*Config, error) {
	var config Config

	if path != "" {
		if err := cleanenv.ReadConfig(path, &config); err != nil {
			return nil, err
		}
	}

	return &config, nil
}
