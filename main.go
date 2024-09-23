package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/call-me-przemo/go-todo/config"
	"github.com/call-me-przemo/go-todo/routes"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := config.GetEnv("APP_PORT")
	fmt.Printf("Server listening at: http://0.0.0.0:%s", port)
	err := http.ListenAndServe(":"+port, routes.Init())

	if err != nil {
		config.Logger.Log(err)
		log.Fatal(err)
	}
}
