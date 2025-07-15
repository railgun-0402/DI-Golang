package domain

import (
	"errors"
	"time"

	"github.com/oklog/ulid/v2"
)

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

func newOrder(
	productID string,
	userID string,
	totalAmount int64,
	products []OrderProduct,
	orderedAt time.Time,
) (*Order, error){
	// productIDのバリデーション
	_, err := ulid.Parse(productID)
	if err != nil {
		return nil, errors.New("商品IDの値が不正です。")
	}

	// 購入金額のバリデーション
	if totalAmount < 0 {
		return nil, errors.New("購入金額の値が不正です。")
	}

	// 購入商品のバリデーション
	if len(products) < 1 {
		return nil, errors.New("購入商品がありません。")
	}

	return &Order{
		ID: productID,
		UserID: userID,
		TotalAmount: totalAmount,
		Products: products,
		orderedAt: orderedAt,
	}, nil
}

func NewOrderProduct(productID string, price int64, quantity int) (*OrderProduct, error) {
	// 商品IDのバリデーション
	_, err := ulid.Parse(productID)
	if err != nil {
		return nil, errors.New("商品IDの値が不正です。")
	}

	// 購入数のバリデーション
	if quantity < 1 {
		return nil, errors.New("購入数の値が不正です。")
	}

	return &OrderProduct{
		productID: productID,
		price:     price,
		quantity:  quantity,
	}, nil
}

func NewOrder(
	productID string,
	userID string,
	totalAmount int64,
	products []OrderProduct,
	now time.Time,
) (*Order, error) {
	return newOrder(
		ulid.Make().String(),
		userID,
		totalAmount,
		products,
		now,
	)
}
