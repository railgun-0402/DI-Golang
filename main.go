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


func main() {
	cfg, _ := config.LoadDefaultConfig(context.TODO())
	db := dynamodb.NewFromConfig(cfg)

	tableName := os.Getenv("DEVICE_TABLE_NAME")

	repo := device.NewDynamoDeviceRepository(db, tableName)
	uc := &usecase.DeviceUsecase{Repo: repo}
	h := &handler.DeviceHandler{Usecase: uc}

	e := echo.New()
	e.POST("/device", h.Register)

	e.Logger.Fatal(e.Start(":8080"))
}

// func product() {
// 	repo := infra.NewMemory()
// 	uc := usecase.NewProductUsecase(repo)

// 	p, err := uc.Save("sample001", "Sample Product", "This is Sample Product", 1000, 15)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Registered product: %+v\n", p)

// 	found, err := repo.FindByID(p.ID)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Found Product: %+v\n", found)
// }

func cognito() {
	http.HandleFunc("/", cog.HandleHome)
	http.HandleFunc("/login", cog.HandleLogin)
	http.HandleFunc("/logout", cog.HandleLogout)
	http.HandleFunc("/callback", cog.HandleCallback)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}