package translation

import (
	"context"

	"translation/internal/svc"
	"translation/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLangListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLangListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLangListLogic {
	return &GetLangListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLangListLogic) GetLangList(req *types.GetLangListReq) (resp *types.GetLangListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
