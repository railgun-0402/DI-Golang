package main

import (
	"log"

	"github.com/labstack/echo/v4"
)


func main() {
	// Echoインスタンス作成
	e := echo.New()

	log.Println("Server running on :8080")
	// log.Fatal(http.ListenAndServe(":8080", handler))
	log.Fatal(e.Start(":8080"))
}