package lang

import (
	"context"
	"translation/internal/svc"
	"translation/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPlatformListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPlatformListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPlatformListLogic {
	return &GetPlatformListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPlatformListLogic) GetPlatformList() (resp *types.GetPlatformListResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.GetPlatformListResp{
		Platformist: make([]types.PlatformData, 0),
	}
	for _, v := range l.svcCtx.LangInfo.TypeList {
		resp.Platformist = append(resp.Platformist, types.PlatformData{
			Name:    v.Name,
			Code:    v.Code,
			Default: v.IsDefault,
		})
	}
	return resp, nil
}
