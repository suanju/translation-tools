package translation

import (
	"net/http"
	"translation/internal/utils/errorx"
	"translation/internal/utils/results"

	"github.com/zeromicro/go-zero/rest/httpx"
	zeroMicroHttp "github.com/zeromicro/x/http"
	"translation/internal/logic/translation"
	"translation/internal/svc"
	"translation/internal/types"
)

func JsonTranslationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TranslationJsonReq
		if err := httpx.Parse(r, &req); err != nil {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, errorx.NewCodeError(results.CodeTypeError, err))
			return
		}
		if validateErr := svcCtx.Validate(&req); validateErr != nil {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, errorx.NewCodeError(results.CodeTypeError, validateErr))
			return
		}

		l := translation.NewJsonTranslationLogic(r.Context(), svcCtx)
		resp, err := l.JsonTranslation(&req)
		if err != nil {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
