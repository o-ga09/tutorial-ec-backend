package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env Env
	App App
	Server Server
	DB DBConfig
	ReadDB ReadDBConfig
	Redis Redis
}

type Env struct {
	Env string `env:"ENV" json:"env,omitempty"`
}

type App struct {
	ProjectID string `env:"PROJECTID" json:"project_id,omitempty"`
}

type Server struct {
	Address string `env:"ADDRESS" json:"host,omitempty"`
	Port string `env:"PORT" json:"port,omitempty"`
}

type DBConfig struct {
	DB_URL string `env:"DB_URL" json:"db___url,omitempty"`
}

type ReadDBConfig struct {
	DB_URL string `env:"DB_URL" json:"db___url,omitempty"`
}

type Redis struct {
	Host string `env:"HOST" json:"host,omitempty"`
	Port string `env:"PORT" json:"port,omitempty"`
}

var (
	once sync.Once
	config Config
)

func GetConfig() *Config {
	once.Do(
		func() {
			if err := envconfig.Process("",&config); err != nil {
				panic(err)
			}
		})
	return &config
}