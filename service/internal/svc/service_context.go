package svc

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/zeromicro/go-zero/rest"
	"reflect"
	"translation/internal/config"
	"translation/internal/middleware"
)

type TypeList []struct {
	Name      string `json:"name"`
	Code      string `json:"code"`
	IsDefault bool   `json:"default"`
}
type LangList map[string]map[string]struct {
	Code     string `json:"code"`
	Original bool   `json:"original"`
	Results  bool   `json:"results"`
}
type LangInfo struct {
	TypeList `json:"type"`
	LangList `json:"lang_list"`
}

type ServiceContext struct {
	Config config.Config
	LangInfo
	TranslationMiddleware rest.Middleware
}

func NewServiceContext(c config.Config, list LangInfo) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		LangInfo:              list,
		TranslationMiddleware: middleware.NewTranslationMiddleware().Handle,
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
