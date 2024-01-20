package utils

import (
	"errors"
	"os"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		panic(errors.New(key + " Environment variable not set"))
	}
	return value
}
