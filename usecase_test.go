package usecase

import (
	"github.com/golang/mock/gomock"
	islib "github.com/matryer/is"
	"github.com/pkg/errors"
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
		{
			name: "user_service_fail_to_find_user_by_ID",
			fieldsFunc: func(t *testing.T, ctrl *gomock.Controller) fields {
				t.Helper()
				userSvc := mock_usecases.NewMockUserService(ctrl)
				userSvc.EXPECT().Get(gomock.Any()).Return(model.User{}, errors.New("unable to find a user")).Times(1)
				return fields{
					userSvc:    userSvc,
					productSvc: mock_usecases.NewMockProductService(ctrl),
				}
			},
			args: args{
				userID: "some-user-ID",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "unable_to_get_products_for_a_user",
			fieldsFunc: func(t *testing.T, ctrl *gomock.Controller) fields {
				t.Helper()
				userSvc := mock_usecases.NewMockUserService(ctrl)
				userSvc.EXPECT().Get(gomock.Any()).Return(model.User{
					ID:   "some-user-ID",
					Name: "name",
				}, nil).Times(1)
				productSvc := mock_usecases.NewMockProductService(ctrl)
				productSvc.EXPECT().GetProducts(gomock.Any()).Return(nil, errors.New("get products")).Times(1)
				return fields{
					userSvc:    userSvc,
					productSvc: productSvc,
				}
			},
			args: args{
				userID: "some-user-ID",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			is := islib.New(t)
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
