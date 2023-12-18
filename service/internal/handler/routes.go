// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	lang "translation/internal/handler/lang"
	translation "translation/internal/handler/translation"
	"translation/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/get_lang_list",
				Handler: lang.GetLangListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/get_platform_list",
				Handler: lang.GetPlatformListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/lang"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.TranslationMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/translation_json",
					Handler: translation.JsonTranslationHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/translation"),
		rest.WithTimeout(30000*time.Millisecond),
	)
}
