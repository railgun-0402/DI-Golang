package usecase_test

import (
	"context"
	"testing"

	domain "github.com/railgun-0402/DI-Golang/app/domain/notification"
	usecase "github.com/railgun-0402/DI-Golang/app/usecase/notification"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockQueue struct {
	mock.Mock
}

// Enqueue usecase層のキュー取得
func (m *MockQueue) Enqueue(ctx context.Context, message string) error {
	args := m.Called(ctx, message)
	return args.Error(0)
}

// ReceiveMessages usecase層メッセージ取得
func (m *MockQueue) ReceiveMessages(ctx context.Context) ([]domain.Message, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Message), args.Error(1)
}

// DeleteMessage usecase層のキュー削除
func (m *MockQueue) DeleteMessage(ctx context.Context, receiptHandle string) error {
	args := m.Called(ctx, receiptHandle)
	return args.Error(0)
}

// Save usecase層DyanamoDBへデータ追加
func (m *MockQueue) Save(ctx context.Context, n domain.Notification) error {
	args := m.Called(ctx, n)
	return args.Error(0)
}

func TestNotificationUsecase_Enqueue(t *testing.T) {
	ctx := context.Background()

	mockQueue := new(MockQueue)
	mockRepo := new(MockQueue)

	uc := usecase.NewNotificationUsecase(mockQueue, mockRepo)

	userID := "u123"
	title := "Test Title"
	message := "Test Message"

	// 事前呼び出し
	mockQueue.
		On("Enqueue", mock.Anything, mock.AnythingOfType("string")).
		Return(nil).
		Once()

	err := uc.Enqueue(ctx, userID, title, message)

	assert.NoError(t, err)
	mockQueue.AssertExpectations(t)
}
