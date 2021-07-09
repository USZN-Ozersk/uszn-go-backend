package config

import (
	"github.com/BurntSushi/toml"
)

// Config ...
type Config struct {
	BindAddr     string `toml:"bind_addr"`
	LogLevel     string `toml:"log_level"`
	DB_URL       string `toml:"db_url"`
	KeyString    string `toml:"key_string"`
	Admin        string `toml:"user"`
	Password     string `toml:"password"`
	UseSSL       bool   `toml:"use_ssl"`
	MailServer   string `toml:"mail_server"`
	MailPort     int    `toml:"mail_port"`
	MailUser     string `toml:"mail_user"`
	MailPassword string `toml:"mail_password"`
}

// New ...
func New() *Config {
	return &Config{
		BindAddr:     ":8080",
		LogLevel:     "debug",
		DB_URL:       "postgres://api:password@pgdb/api?sslmode=disable",
		KeyString:    "passphrase",
		Admin:        "admin",
		Password:     "password",
		UseSSL:       false,
		MailServer:   "127.0.0.1",
		MailPort:     25,
		MailUser:     "User",
		MailPassword: "Password",
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
