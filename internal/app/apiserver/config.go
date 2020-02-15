package apiserver

import (
	"github.com/kapustaprusta/simple_rest_server/internal/app/store"
)

// Config ...
type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LogLevel    string `toml:"log_level"`
	StoreConfig *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":8080",
		LogLevel:    "debug",
		StoreConfig: store.NewConfig(),
	}
}
