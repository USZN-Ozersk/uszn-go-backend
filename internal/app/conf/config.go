package conf

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

// New ...
func New() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}

// InitConfig ...
func (config *Config) InitConfig(configpath string) error {
	_, err := toml.DecodeFile(configpath, config)
	if err != nil {
		return err
	}

	return nil
}
