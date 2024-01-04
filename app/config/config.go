package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	Env          string `env:"ENV" envDefault:"dev"`
	Port         string `env:"PORT" envDefault:"80"`
	Database_url string `env:"DB_URL" envDefault:""`
	RedisURL string `env:"REDIS_URL" envDefault:""`
	ProjectID    string `env:"PROJECTID" envDefault:""`
}
// type Config struct {
// 	Env Env
// 	App App
// 	Server Server
// 	DB DBConfig
// 	ReadDB ReadDBConfig
// 	Redis Redis
// }

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
	DB_URL string `env:"DB_URL" json:"db_url,omitempty"`
}

type ReadDBConfig struct {
	DB_URL string `env:"DB_URL" json:"db_url,omitempty"`
}

type Redis struct {
	Host string `env:"HOST" json:"host,omitempty"`
	Port string `env:"PORT" json:"port,omitempty"`
}

func GetConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil
	}
	return cfg
}