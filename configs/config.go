package configs

import (
	"log"
	"syscall"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
}

func loadEnv() {

	err := godotenv.Overload(ENV_FILE)
	if err != nil {
		log.Fatalf("error in loading .env file:%s", err.Error())
	}
}

func GetEnvWithKey(key, defaultValue string) string {
	value,isFound:= syscall.Getenv(key)
	if !isFound{
		syscall.Setenv(key,defaultValue)
		return defaultValue
	}
	return value
}
