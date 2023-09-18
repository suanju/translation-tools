package errorx

import (
	"github.com/zeromicro/x/errors"
	"translation/api/internal/utils/results"
)

func NewCodeError(code int, err error) error {
	return &errors.CodeMsg{
		Code: code,
		Msg:  err.Error(),
	}
}

func NewDefaultError(err error) error {
	return NewCodeError(results.CodeDefaultError, err)
}
