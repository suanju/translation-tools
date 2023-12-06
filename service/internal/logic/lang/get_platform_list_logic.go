package lang

import (
	"context"
	"fmt"

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
	fmt.Println(123123)
	return
}
