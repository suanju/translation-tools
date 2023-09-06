package login

import (
	"net/http"
	"translation/api/internal/utils/errorx"
	"translation/api/internal/utils/results"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	zeromicrohttp "github.com/zeromicro/x/http"
	"translation/api/internal/logic/users/login"
	"translation/api/internal/svc"
	"translation/api/internal/types"
)

func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			zeromicrohttp.JsonBaseResponseCtx(r.Context(), w, errorx.NewCodeError(results.CodeTypeError, err))
			return
		}

		l := login.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		if err != nil {
			zeromicrohttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			zeromicrohttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
