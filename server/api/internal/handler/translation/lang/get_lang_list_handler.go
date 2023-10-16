package lang

import (
	"net/http"
	"translation/api/internal/utils/errorx"
	"translation/api/internal/utils/results"

	"github.com/zeromicro/go-zero/rest/httpx"
	zeroMicroHttp "github.com/zeromicro/x/http"
	"translation/api/internal/logic/translation/lang"
	"translation/api/internal/svc"
	"translation/api/internal/types"
)

func GetLangListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetLangListReq
		if err := httpx.Parse(r, &req); err != nil {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, errorx.NewCodeError(results.CodeTypeError, err))
			return
		}
		if validateErr := svcCtx.Validate(&req); validateErr != nil {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, errorx.NewCodeError(results.CodeTypeError, validateErr))
			return
		}

		l := lang.NewGetLangListLogic(r.Context(), svcCtx)
		resp, err := l.GetLangList(&req)
		if err != nil {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
