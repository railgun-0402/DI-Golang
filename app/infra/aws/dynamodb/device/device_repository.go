package device

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	domain "github.com/railgun-0402/DI-Golang/app/domain/device"
)

type DynamoDeviceRepository struct {
    Client    *dynamodb.Client
    TableName string
}

func NewDynamoDeviceRepository(c *dynamodb.Client, table string) domain.DeviceRepository {
	return &DynamoDeviceRepository{Client: c, TableName: table}
}

func (r *DynamoDeviceRepository) Save(device domain.Device) error {
	item, err := attributevalue.MarshalMap(device)
	if err != nil {
		return err
	}

	_, err = r.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item: item,
	})
	return err
}