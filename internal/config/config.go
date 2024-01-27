package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Env         string `env:"ENV" env-default:"local"`
	Port        int    `env:"PORT" env-default:"8080" required:"true"`
	DatabaseURL string `env:"DATABASE_URL" required:"true"`
}

func Load() *Config {
	path := "./.env"
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}
	return &cfg
}
