# go-table-testing-arg-functions

Example on how to use fields/args generators in table-tests

# Generic Pattern to build table-driven tests with args/field "factories":

1. Generate table-driven test stub, like this:
```go
func TestProductsForUser_Get(t *testing.T) {
    type fields struct {
        userSvc    UserService
        productSvc ProductService
    }
    type args struct {
        userID model.UserID
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    []model.Product
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        tt := tt
        t.Run(tt.name, func(t *testing.T) {
            pu := &ProductsForUser{
                userSvc:    tt.fields.userSvc,
                productSvc: tt.fields.productSvc,
            }
            got, err := pu.Get(tt.args.userID)
            if (err != nil) != tt.wantErr {
                t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("Get() got = %v, want %v", got, tt.want)
            }
        })
    }
}
```

**NOTE:** And do not forget about assignments `tt := tt` in case those will be converted to parallel tests in the future (using `t.Parallel()`):
```go
    for _, tt := range tests {
        tt := tt // <-- THIS IS IMPORTANT
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            // write actual test			
        })
    }

```

2. Change the `args` and/or `fields` parameter to be a function that returns filled `args`/`fields` respectively:
```go
    type fields struct {
        userSvc    UserService
        productSvc ProductService
    }
    type args struct {
        userID model.UserID
    }
    tests := []struct {
        name       string
        fieldsFunc func(t *testing.T, ctrl *gomock.Controller) fields // <-- THIS IS a "factory"
        args       args // <-- In some cases, this also can be converted to a "factory" 
        want       []model.Product
        wantErr    bool
    }{
        // TODO: Add test cases.
    }
```

Typical "factory" function signature may look like:
```go
func(t *testing.T) fields
```
or in case of [GoMock](https://github.com/golang/mock)
```go
func(t *testing.T, ctrl *gomock.Controller) fields
```

3. The benefits of using this "factory" approach to prepare all the fields inside `args`/`fields` are:
    - the ability to code custom preparation logic;
    - the ability to build one mock/stub that is dependent on another one, i.e., introduce coupling between components;

As shown in the example:
```go
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
}
```