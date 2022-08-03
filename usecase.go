package usecase

import (
	"github.com/pkg/errors"
	"go-table-testing-arg-functions/model"
)

//go:generate mockgen -destination=./mocks/usecase_mock.go -package=mock_usecases . UserService,ProductService

// UserService is a User service.
type UserService interface {
	Get(id model.UserID) (model.User, error)
}

// ProductService is a Product service.
type ProductService interface {
	GetProducts(userName string) ([]model.Product, error)
}

// ProductsForUser is a "get products for a user" use-case.
type ProductsForUser struct {
	userSvc    UserService
	productSvc ProductService
}

// NewProductsForUser creates new ProductsForUser use-case.
func NewProductsForUser(userSvc UserService, productSvc ProductService) *ProductsForUser {
	return &ProductsForUser{
		userSvc:    userSvc,
		productSvc: productSvc,
	}
}

// Get gets products based on username.
func (pu *ProductsForUser) Get(userID model.UserID) ([]model.Product, error) {
	if userID == "" {
		return nil, errors.New("user ID can't be empty")
	}
	user, err := pu.userSvc.Get(userID)
	if err != nil {
		return nil, errors.WithMessagef(err, "user with id=%q", userID)
	}
	products, err := pu.productSvc.GetProducts(user.Name)
	if err != nil {
		return nil, errors.WithMessagef(err, "get products by username=%q", user.Name)
	}
	return products, nil
}
