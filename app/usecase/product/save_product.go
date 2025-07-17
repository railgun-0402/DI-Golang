package usecase

import domain "github.com/railgun-0402/DI-Golang/app/domain/product"

type ProductUsecase struct {
	Repo domain.ProductRepositoy
}

func NewProductUsecase(repo domain.ProductRepositoy) *ProductUsecase {
	return &ProductUsecase{Repo: repo}
}

func (uc *ProductUsecase) Save(id, name, description string, price int64, stock int) (*domain.Product, error) {
	p, err := domain.NewProduct(id, name, description, price, stock)
	if err != nil {
		return nil, err
	}

	err = uc.Repo.Save(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

