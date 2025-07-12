package infra

import (
	"errors"
	"sync"

	domain "github.com/railgun-0402/DI-Golang/app/domain/product"
)

// TODO: 今日は眠いのでメモリの保存してるが、最終的にDynamoDBを使用する
type Memory struct {
	mu sync.RWMutex
	products map[string]*domain.Product
}

func NewMemory() *Memory {
	return &Memory{
		products: make(map[string]*domain.Product),
	}
}

func (r *Memory) Save(p *domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.products[p.ID]; exists {
		return errors.New("既に登録されてる商品です！")
	}
	r.products[p.ID] = p
	return nil
}

func (r *Memory) FindByID(id string) (*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	p, ok := r.products[id]
	if !ok {
		return nil, errors.New("商品が見つかりません！")
	}

	return p, nil
}
