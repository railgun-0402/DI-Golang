package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/labstack/echo/v4"
	"github.com/railgun-0402/DI-Golang/app/handler"
	cog "github.com/railgun-0402/DI-Golang/app/infra/aws/cognito"
	"github.com/railgun-0402/DI-Golang/app/infra/aws/dynamodb/device"
	usecase "github.com/railgun-0402/DI-Golang/app/usecase/device"
)

// TODO: 早くhandlerごとに処理をまとめてmain.goを綺麗にしてえ(現状やばすぎ)
func main() {
	ctx := context.Background()
	cfg, _ := config.LoadDefaultConfig(ctx)
	db := dynamodb.NewFromConfig(cfg)

	tableName := os.Getenv("DEVICE_TABLE_NAME")

	repo := device.NewDynamoDeviceRepository(db, tableName)
	uc := &usecase.DeviceUsecase{Repo: repo}
	h := &handler.DeviceHandler{Usecase: uc}

	e := echo.New()
	e.POST("/device", h.Register)

	e.Logger.Fatal(e.Start(":8080"))
}


func cognito() {
	http.HandleFunc("/", cog.HandleHome)
	http.HandleFunc("/login", cog.HandleLogin)
	http.HandleFunc("/logout", cog.HandleLogout)
	http.HandleFunc("/callback", cog.HandleCallback)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}