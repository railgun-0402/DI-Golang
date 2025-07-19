package notification

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	domain "github.com/railgun-0402/DI-Golang/app/domain/notification"
)

type DynamoNotificationRepository struct {
	Client *dynamodb.Client
	TableName string
}

func NewDynamoNotificationRepository(c *dynamodb.Client, tableName string) *DynamoNotificationRepository {
	return &DynamoNotificationRepository{
		Client: c,
		TableName: tableName,
	}
}

func (r *DynamoNotificationRepository) Save(ctx context.Context, n domain.Notification) error {
	item, err := attributevalue.MarshalMap(n)
	if err != nil {
		return err
	}

	_, err = r.Client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item: item,
	})
	return err
}