package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/railgun-0402/DI-Golang/app/infra"
	cog "github.com/railgun-0402/DI-Golang/app/infra/aws/cognito"
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


func cognito() {
	http.HandleFunc("/", cog.HandleHome)
	http.HandleFunc("/login", cog.HandleLogin)
	http.HandleFunc("/logout", cog.HandleLogout)
	http.HandleFunc("/callback", cog.HandleCallback)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}