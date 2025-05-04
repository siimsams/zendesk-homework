package env

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port      string `env:"PORT" envDefault:"50051"`
	DbPath    string `env:"DB_PATH" envDefault:"database.db"`
	JwtSecret string `env:"JWT_SECRET" envDefault:"your-secret-key"`
}

var config Config

func init() {
	godotenv.Load()

	if err := env.Parse(&config); err != nil {
		panic(err)
	}
}

func GetConfig() Config {
	return config
}

func JwtSecretBytes() []byte {
	return []byte(config.JwtSecret)
}
