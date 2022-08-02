package main

import (
	"go-table-testing-arg-functions/model"
	"go-table-testing-arg-functions/service"
	"go-table-testing-arg-functions/usecase"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	userSvc := service.NewUser()
	productSvc := service.NewProduct()
	productsForUser := usecase.NewProductsForUser(userSvc, productSvc)

	userID := model.UserID("777")
	products, err := productsForUser.Get(userID)
	if err != nil {
		log.Fatalf("failed to get products for a user: %+v", err)
	}
	log.Println("Fetched products for a userID=", userID)
	for _, product := range products {
		log.Println(product)
	}
}
