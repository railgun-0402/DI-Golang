package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/labstack/echo/v4"
	"github.com/railgun-0402/DI-Golang/app/handler"
	cog "github.com/railgun-0402/DI-Golang/app/infra/aws/cognito"
	notification_repository "github.com/railgun-0402/DI-Golang/app/infra/aws/dynamodb/notification"
	notifier "github.com/railgun-0402/DI-Golang/app/infra/aws/sns"
	sqs_repository "github.com/railgun-0402/DI-Golang/app/infra/aws/sqs/notification"

	usecase "github.com/railgun-0402/DI-Golang/app/usecase/notification"
)

// func New(addr string) {
// 	cfg, _ := config.LoadDefaultConfig(context.Background())
// 	queueURL := os.Getenv("SQS_QUEUE_URL")

// 	sqsClient := sqs.NewFromConfig(cfg)
// 	queue := sqs_repository.NewNotificationQueue(sqsClient, queueURL)

// 	uc := usecase.NewNotificationUsecase(queue)
// 	h := handler.NewNotificationHandler(uc)

// 	e := echo.New()
// 	e.POST("/notifications", h.Enqueue)

// 	log.Fatal(http.ListenAndServe(addr, nil))
// }

func cognito() {
	http.HandleFunc("/", cog.HandleHome)
	http.HandleFunc("/login", cog.HandleLogin)
	http.HandleFunc("/logout", cog.HandleLogout)
	http.HandleFunc("/callback", cog.HandleCallback)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func RunWorker(ctx context.Context) {
	cfg, _ := config.LoadDefaultConfig(ctx)

	queueURL := os.Getenv("SQS_QUEUE_URL")
	topicARN := os.Getenv("SNS_TOPIC_ARN")
	notifTable := os.Getenv("NOTIFICATION_TABLE_NAME")

	sqsClient := sqs.NewFromConfig(cfg)
	snsClient := sns.NewFromConfig(cfg)
	dynamodbCLient := dynamodb.NewFromConfig(cfg)

	q := sqs_repository.NewQueue(sqsClient, queueURL)
	n := notifier.NewNotifier(snsClient, topicARN)
	r := notification_repository.NewDynamoNotificationRepository(dynamodbCLient, notifTable)

	uc := usecase.NewWorkerUsecase(q, n, r)
	h := handler.NewWorkerHandler(uc)

	h.Run(ctx)
}

func NewAPI() *echo.Echo {
	ctx := context.Background()
	cfg, _ := config.LoadDefaultConfig(ctx)
	db := dynamodb.NewFromConfig(cfg)

	tableName := os.Getenv("DEVICE_TABLE_NAME")
	queueURL := os.Getenv("SQS_QUEUE_URL")

	deviceRepo := notification_repository.NewDynamoNotificationRepository(db, tableName)
	sqsClient := sqs.NewFromConfig(cfg)
	queue := sqs_repository.NewQueue(sqsClient, queueURL)

	deviceUC := usecase.NewNotificationUsecase(queue, deviceRepo)
	deviceHandler := handler.NewNotificationHandler(deviceUC)

	e := echo.New()
	e.POST("/notifications", deviceHandler.Enqueue)

	return e
}