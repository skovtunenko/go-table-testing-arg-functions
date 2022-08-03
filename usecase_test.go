package usecase

import (
	"github.com/matryer/is"
	"testing"
)

func TestNewProductsForUser(t *testing.T) {
	type args struct {
		userSvc    UserService
		productSvc ProductService
	}
	tests := []struct {
		name string
		args args
		want *ProductsForUser
	}{
		{
			name: "all_dependencies_are_nil",
			args: args{
				userSvc:    nil,
				productSvc: nil,
			},
			want: &ProductsForUser{
				userSvc:    nil,
				productSvc: nil,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := NewProductsForUser(tt.args.userSvc, tt.args.productSvc)
			is.Equal(got, tt.want)
		})
	}
}
