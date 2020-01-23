package config

import (
	"github.com/BurntSushi/toml"
)

// Config ...
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	DB_URL   string `toml:"db_url"`
}

// New ...
func New() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		DB_URL:   "postgres://api:password@pgdb/api?sslmode=disable",
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