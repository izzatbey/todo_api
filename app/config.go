package app

import (
	"github.com/joho/godotenv"
)

func GetEnvVariable(key string) string{
	err := godotenv.Load(".env")
}