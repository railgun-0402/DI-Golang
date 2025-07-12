package product

import (
	"errors"
	"unicode/utf8"

	"github.com/oklog/ulid/v2"
)

type Product struct {
	id string
	ownerID string
	name string
	description string
	price int64
	stock int
}

const (
	MinNameLength = 1
	MaxNameLength = 100

	MinDescriptionLength = 1
	MaxDescriptionLength = 1000
)

func isValid(s string) bool {
	_, err := ulid.Parse(s)
	return err == nil
}

func NewProduct (
	id string,
	ownerID string,
	name string,
	description string,
	price int64,
	stock int,
) (*Product, error) {
	// ownerIDバリデーション
	if !isValid(ownerID) {
		return nil, errors.New("ownerIDが不正です")
	}

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
		id: id,
		ownerID: ownerID,
		name: name,
		description: description,
		price: price,
		stock: stock,
	}, nil
}
