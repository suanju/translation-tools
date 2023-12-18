package middleware

import (
	"context"
	zeroMicroHttp "github.com/zeromicro/x/http"
	"net/http"
	"translation/internal/utils/errorx"
	"translation/internal/utils/results"
)

type TranslationMiddleware struct {
}

func NewTranslationMiddleware() *TranslationMiddleware {
	return &TranslationMiddleware{}
}

func (m *TranslationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//判断是否保存了设置信息
		var platformCode string
		var APPID string
		var KEY string
		if len(r.Header["Platform_code"]) == 0 {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, errorx.NewCodeErrorByDefaultMsg(results.CodeNoPlatformCode))
			return
		}
		platformCode = r.Header["Platform_code"][0]
		r = r.WithContext(context.WithValue(r.Context(), "platform_code", platformCode))
		switch platformCode {
		case "baidu":
			if len(r.Header["App_id"]) == 0 || len(r.Header["App_key"]) == 0 {
				zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, errorx.NewCodeErrorByDefaultMsg(results.CodeNoFillConfig))
				return
			}
			APPID = r.Header["App_id"][0]
			KEY = r.Header["App_key"][0]
			r = r.WithContext(context.WithValue(r.Context(), "app_id", APPID))
			r = r.WithContext(context.WithValue(r.Context(), "app_key", KEY))
			break
		case "deepl":
			if len(r.Header["App_key"]) == 0 {
				zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, errorx.NewCodeErrorByDefaultMsg(results.CodeNoFillConfig))
				return
			}
			KEY = r.Header["App_key"][0]
			r = r.WithContext(context.WithValue(r.Context(), "app_key", KEY))
			break
		default:
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, errorx.NewCodeErrorByDefaultMsg(results.CodeErrorPlatformCode))
			return
		}
		next(w, r)
	}
}
