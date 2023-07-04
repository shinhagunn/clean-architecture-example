package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		DB `yaml:"db"`
	}

	DB struct {
		Host     string `env:"DB_HOST" env-default:"localhost"`
		User     string `env:"DB_USER" env-default:"postgres"`
		Password string `env:"DB_PASSWORD" env-default:"postgres"`
		Name     string `env:"DB_NAME" env-default:"todo"`
		Port     string `env:"DB_PORT" env-default:"5432"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
