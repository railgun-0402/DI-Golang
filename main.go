package main

import (
	"fmt"

	"github.com/railgun-0402/DI-Golang/app/infra"
	usecase "github.com/railgun-0402/DI-Golang/app/usecase/product"
)


func main() {
	repo := infra.NewMemory()
	uc := usecase.NewProductUsecase(repo)

	p, err := uc.Save("sample001", "Sample Product", "This is Sample Product", 1000, 15)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Registered product: %+v\n", p)

	found, err := repo.FindByID(p.ID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Found Product: %+v\n", found)
}