package domain

type ProductRepositoy interface {
	Save(product *Product) error
	FindByID(id string) (*Product, error)
}

