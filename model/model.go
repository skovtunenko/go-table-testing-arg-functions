package model

import "fmt"

type UserID string

type ProductID string

type Product struct {
	ID    ProductID
	Owner string
	Name  string
}

var _ fmt.Stringer = Product{}

func (p Product) String() string {
	return fmt.Sprintf("ProductID=%q, Name=%q, Owner=%q", p.ID, p.Name, p.Owner)
}

type User struct {
	ID   UserID
	Name string
}
