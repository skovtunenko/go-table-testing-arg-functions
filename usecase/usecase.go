package usecase

import (
	"github.com/pkg/errors"
	"go-table-testing-arg-functions/model"
)

type userService interface {
	Get(id model.UserID) (model.User, error)
}

type productService interface {
	GetProducts(userName string) ([]model.Product, error)
}

// ProductsForUser is a "get products for a user" use-case.
type ProductsForUser struct {
	userSvc    userService
	productSvc productService
}

func NewProductsForUser(userSvc userService, productSvc productService) *ProductsForUser {
	return &ProductsForUser{
		userSvc:    userSvc,
		productSvc: productSvc,
	}
}

func (pu *ProductsForUser) Get(userID model.UserID) ([]model.Product, error) {
	user, err := pu.userSvc.Get(userID)
	if err != nil {
		return nil, errors.WithMessagef(err, "no user with id=%q", userID)
	}
	products, err := pu.productSvc.GetProducts(user.Name)
	if err != nil {
		return nil, errors.WithMessagef(err, "get products by username=%q", user.Name)
	}
	return products, nil
}
