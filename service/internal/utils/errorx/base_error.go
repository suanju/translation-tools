package errorx

import (
	xerror "github.com/zeromicro/x/errors"
	"translation/internal/utils/results"
)

func NewCodeError(code int, err error) error {
	return &xerror.CodeMsg{
		Code: code,
		Msg:  err.Error(),
	}
}

func NewCodeErrorByDefaultMsg(code int) error {
	return &xerror.CodeMsg{
		Code: code,
		Msg:  results.Message[code],
	}
}
