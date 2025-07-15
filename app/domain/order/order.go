package domain

import "time"

type Order struct {
	ID string
	UserID string
	TotalAmount int64
	Products OrderProducts
	orderedAt time.Time
}

type OrderProducts []OrderProduct

type OrderProduct struct {
	productID string
	price int64
	quantity int
}