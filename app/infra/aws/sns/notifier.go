package notifier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type Notifier struct {
	client *sns.Client
	topicARN string
}

func NewNotifier(client *sns.Client, topicARN string) *Notifier {
	return &Notifier{
		client:   client,
		topicARN: topicARN,
	}
}

func (n *Notifier) Notify(ctx context.Context, message string) error {
	_, err := n.client.Publish(ctx, &sns.PublishInput{
		TopicArn: aws.String(n.topicARN),
		Message: aws.String(message),
	})
	return err
}