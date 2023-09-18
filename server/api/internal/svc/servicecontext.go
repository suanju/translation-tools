package svc

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/zrpc"
	"reflect"
	"translation/api/internal/config"
	"translation/api/internal/utils/jwt"
	"translation/rpc/user/userservice"
)

type ServiceContext struct {
	Config      config.Config
	AuthService *jwt.AuthService
	UserService userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AuthService: jwt.NewAuthService(c.Auth.AccessSecret, c.Auth.AccessExpire),
		UserService: userservice.NewUserService(zrpc.MustNewClient(c.User)),
	}
}

func (svc ServiceContext) Validate(dataStruct interface{}) error {
	zhCh := zh.New()
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("label")
		return name
	})
	uni := ut.New(zhCh)
	trans, _ := uni.GetTranslator("zh")
	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(dataStruct)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}
