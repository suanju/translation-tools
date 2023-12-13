package lang

import (
	zeroMicroHttp "github.com/zeromicro/x/http"
	"net/http"
	"translation/internal/logic/lang"
	"translation/internal/svc"
)

func GetPlatformListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := lang.NewGetPlatformListLogic(r.Context(), svcCtx)
		resp, err := l.GetPlatformList()
		if err != nil {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			zeroMicroHttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
