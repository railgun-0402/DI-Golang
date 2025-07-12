package domain

import (
	"errors"
	"unicode/utf8"

	"github.com/oklog/ulid/v2"
)

type Product struct {
	ID string
	Name string
	Description string
	Price int64
	Stock int
}

const (
	MinNameLength = 1
	MaxNameLength = 100

	MinDescriptionLength = 1
	MaxDescriptionLength = 1000
)

// TODO: テストコードを実装する
func newProduct (
	id string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	if utf8.RuneCountInString(name) < MinNameLength {
		return nil, errors.New("商品名を入力してください。")
	}

	if utf8.RuneCountInString(name) > MaxNameLength {
		return nil, errors.New("商品名は100文字以下で入力してください。")
	}

	if utf8.RuneCountInString(description) < MinDescriptionLength {
		return nil, errors.New("商品の説明を入力してください。")
	}

	if utf8.RuneCountInString(description) > MaxDescriptionLength {
		return nil, errors.New("商品説明は1000文字以下で入力してください。")
	}

	// 価格のバリデーション
	if price < 1 {
		return nil, errors.New("価格の値が不正です。")
	}

	if stock < 0 {
		return nil, errors.New("在庫数の値が不正です。")
	}
	return &Product{
		ID: id,
		Name: name,
		Description: description,
		Price: price,
		Stock: stock,
	}, nil
}

func Reconstruct(
	id string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	return newProduct(
		id,
		name,
		description,
		price,
		stock,
	)
}

func NewProduct(
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	return newProduct(
		ulid.Make().String(),
		name,
		description,
		price,
		stock,
	)
}
