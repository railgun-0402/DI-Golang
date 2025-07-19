package usecase

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	domain "github.com/railgun-0402/DI-Golang/app/domain/notification"
)

type WorkerUsecase struct {
	queue domain.Queue
	notifier domain.Notifier
	repository domain.NotificationRepository
}

func NewWorkerUsecase(q domain.Queue, n domain.Notifier, r domain.NotificationRepository) *WorkerUsecase {
	return &WorkerUsecase{
		queue:   q,
		notifier:   n,
		repository: r,
	}
}

func (uc *WorkerUsecase) Run(ctx context.Context) {
	for {
		msgs, err := uc.queue.ReceiveMessages(ctx)
		if err != nil {
			log.Println("failed to receive messages:", err)
			time.Sleep(2 * time.Second)
			continue
		}

		for _, msg := range msgs {
			var n domain.Notification
			if err := json.Unmarshal([]byte(msg.Body), &n); err != nil {
				log.Println("invalid message:", err)
				continue
			}

			n.ID = uuid.New().String()
			n.CreatedAt = time.Now()

			if err := uc.notifier.Notify(ctx, n.Message); err != nil {
				log.Println("failed to notify:", err)
			}

			if err := uc.repository.Save(ctx, n); err != nil {
				log.Println("failed to save:", err)
			}

			// if err := uc.queue.DeleteMessage(ctx, msg.ReceiptHandle); err != nil {
			// 	log.Println("failed to delete message:", err)
			// }
		}
	}
}
