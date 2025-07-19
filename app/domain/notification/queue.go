package domain

import (
	"context"
)

type Message struct {
	Body          string
	ReceiptHandle string
}

type Queue interface {
	ReceiveMessages(ctx context.Context) ([]Message, error)
	DeleteMessage(ctx context.Context, receiptHandle string) error
	Enqueue(ctx context.Context, message string) error
}
