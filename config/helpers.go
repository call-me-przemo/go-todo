package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
)

var Logger, _ = thoth.Init("log")

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		Logger.Log(errors.New("no .env file found"))
		log.Fatal("No .env file found")
	}
}

func GetEnv(name string) string {

	env, exist := os.LookupEnv(name)

	if !exist {
		msg := fmt.Sprintf("%s not set in .env", name)
		Logger.Log(errors.New(msg))
		log.Fatal(msg)
	}

	return env
}
