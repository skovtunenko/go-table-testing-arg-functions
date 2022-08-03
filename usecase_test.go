package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/matryer/is"
	mock_usecases "go-table-testing-arg-functions/mocks"
	"go-table-testing-arg-functions/model"
	"testing"
)

func TestProductsForUser_Get(t *testing.T) {
	type fields struct {
		userSvc    UserService
		productSvc ProductService
	}
	type args struct {
		userID model.UserID
	}
	tests := []struct {
		name       string
		fieldsFunc func(t *testing.T, ctrl *gomock.Controller) fields
		args       args
		want       []model.Product
		wantErr    bool
	}{
		{
			name: "empty_userID",
			fieldsFunc: func(t *testing.T, ctrl *gomock.Controller) fields {
				t.Helper()
				return fields{
					userSvc:    mock_usecases.NewMockUserService(ctrl),
					productSvc: mock_usecases.NewMockProductService(ctrl),
				}
			},
			args: args{
				userID: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			ctrl := gomock.NewController(t)

			// get prepared fields data for the current test:
			fields := tt.fieldsFunc(t, ctrl)

			pu := NewProductsForUser(fields.userSvc, fields.productSvc)
			got, err := pu.Get(tt.args.userID)
			if !tt.wantErr {
				is.NoErr(err)
			}
			is.Equal(tt.want, got)
		})
	}
}
