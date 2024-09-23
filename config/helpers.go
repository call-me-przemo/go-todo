package config

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ichtrojan/thoth"
)

var Logger, _ = thoth.Init("log")

func GetEnv(name string) string {

	env, exist := os.LookupEnv(name)

	if !exist {
		msg := fmt.Sprintf("%s not set in .env", name)
		Logger.Log(errors.New(msg))
		log.Fatal(msg)
	}

	return env
}
