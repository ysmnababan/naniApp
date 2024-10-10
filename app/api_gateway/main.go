package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
func main() {
	e := echo.New()
	log.Fatal(e.Start(":" + os.Getenv("GATEWAY_PORT")))
}
