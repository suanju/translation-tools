package translation

import (
	"context"

	"translation/internal/svc"
	"translation/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JsonTranslationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJsonTranslationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JsonTranslationLogic {
	return &JsonTranslationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JsonTranslationLogic) JsonTranslation(req *types.TranslationJsonReq) (resp *types.TranslationJsonResp, err error) {
	// todo: add your logic here and delete this line

	return
}
