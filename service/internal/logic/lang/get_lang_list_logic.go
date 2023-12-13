package lang

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
	resp = &types.GetLangListResp{
		LangList: make([]types.LangData, 0),
	}
	for k, v := range l.svcCtx.LangInfo.LangList {
		for kk, vv := range v {
			if kk == req.Platform {
				resp.LangList = append(resp.LangList, types.LangData{
					Lang:     k,
					Code:     vv.Code,
					Original: vv.Original,
					Results:  vv.Results,
				})
			}
		}
	}
	return resp, err
}
