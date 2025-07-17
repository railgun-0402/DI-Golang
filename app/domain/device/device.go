package domain

import "time"

type Device struct {
	ID        string    `dynamodbav:"id"`
	UserID    string    `dynamodbav:"user_id"`
	Token     string    `dynamodbav:"token"`
	Platform  string    `dynamodbav:"platform"`
	CreatedAt time.Time `dynamodbav:"created_at"`
}

type DeviceRepository interface {
	Save(device Device) error
	// TODO: 後程実装
	// FindByUserID(userID string) ([]Device, error)
	// Delete(id string) error
}

// TODO: 後程実装
// type NotificationQueue interface {
// 	Enqueue(notification Notification) error
// }