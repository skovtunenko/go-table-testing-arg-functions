package service

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"
	"go-table-testing-arg-functions/model"
	"math/rand"
)

// Product is a product service.
type Product struct{}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) GetProducts(userName string) ([]model.Product, error) {
	if userName == "" {
		return nil, errors.New("user name can't be empty")
	}
	// simulate data retrieval process....
	products := make([]model.Product, rand.Int31n(10))
	for i := range products {
		products[i] = model.Product{
			ID:    model.ProductID(gofakeit.UUID()),
			Owner: userName,
			Name:  gofakeit.AppName(),
		}
	}
	return products, nil
}

// User is a user service.
type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) Get(id model.UserID) (model.User, error) {
	if id == "" {
		return model.User{}, errors.New("userID can't be empty")
	}
	// simulate data retrieval process....
	return model.User{
		ID:   id,
		Name: gofakeit.Name(),
	}, nil
}
