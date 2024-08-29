package utils

import (
	"os"

	"github.com/mingi3442/ibc-monitoring/internal/types"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	ChainA types.ConfigParams
	ChainB types.ConfigParams
}

func ReadConfig() (*Config, error) {
	file, err := os.Open("config.toml")

	if err != nil {
		return &Config{}, err
	}
	defer file.Close()

	var config Config
	decoder := toml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return &Config{}, err
	}

	return &config, nil
}
