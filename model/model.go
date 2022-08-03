package model

import "fmt"

type UserID string

type UserName string

type ProductID string

type ProductName string

type Product struct {
	ID    ProductID
	Owner UserName
	Name  ProductName
}

var _ fmt.Stringer = Product{}

func (p Product) String() string {
	return fmt.Sprintf("ProductID=%q, Name=%q, Owner=%q", p.ID, p.Name, p.Owner)
}

type User struct {
	ID   UserID
	Name UserName
}
