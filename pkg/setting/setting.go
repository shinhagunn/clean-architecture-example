package setting

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	DB struct {
		Host     string `yaml:"DB_HOST" env:"DB_HOST" env-default:"localhost"`
		User     string `yaml:"DB_USER" env:"DB_USER" env-default:"postgres"`
		Password string `yaml:"DB_PASSWORD" env:"DB_PASSWORD" env-default:"postgres"`
		Name     string `yaml:"DB_NAME" env:"DB_NAME" env-default:"todo"`
		Port     int64  `yaml:"DB_PORT" env:"DB_PORT" env-default:"5432"`
	} `yaml:"DB"`

	Server struct {
		Port         int `yaml:"SERVER_PORT" env:"SERVER_PORT" env-default:"3000"`
		ReadTimeout  int `yaml:"SERVER_READ_TIMEOUT" env:"SERVER_READ_TIMEOUT" env-default:"60"`
		WriteTimeout int `yaml:"SERVER_WRITE_TIMEOUT" env:"SERVER_WRITE_TIMEOUT" env-default:"60"`
	} `yaml:"SERVER"`

	SecretKey string `yaml:"SECRET_KEY" env:"SECRET_KEY" env-default:"LgszAIV5dBwmy93buOvFDSvmXaam1rzjNAHQ9HJNizE"`
}

var Cfg *config

func Setup() {
	Cfg = &config{}

	if err := cleanenv.ReadConfig("config/config.yml", Cfg); err != nil {
		panic(err)
	}
}
