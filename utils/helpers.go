package utils

import (
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		panic(errors.New(key + " Environment variable not set"))
	}
	return value
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
