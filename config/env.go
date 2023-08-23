package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Env struct {
	Port         string
	Dsn          string
	RedisConn    string
	InProduction bool
}

func GetEnvironment() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Env{
		Port:         os.Getenv("APP_PORT"),
		Dsn:          os.Getenv("DATABASE_URL"),
		RedisConn:    os.Getenv("UPSTASH_REDIS_CONNECTION"),
		InProduction: os.Getenv("APP_ENV") == "production",
	}, nil
}
