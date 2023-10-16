package lang

import (
	"context"
	"fmt"
	"translation/api/internal/utils/errorx"
	"translation/rpc/translation/pb"

	"translation/api/internal/svc"
	"translation/api/internal/types"

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
	langList, err := l.svcCtx.TranslationService.GetLangList(l.ctx, &pb.GetLangListReq{})
	if err != nil {
		fmt.Println(err)
		return nil, errorx.NewDefaultErrorMessage("获取失败")
	}
	list := make([]types.LangData, 0)
	for _, v := range langList.LangList {
		list = append(list, types.LangData{
			Id:       v.Id,
			Lang:     v.Lang,
			Code:     v.Code,
			Original: v.Original,
			Results:  v.Results,
		})
	}
	return &types.GetLangListResp{LangList: list}, err
}
