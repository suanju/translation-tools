package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"translation/rpc/translation/internal/svc"
	"translation/rpc/translation/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLangListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLangListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLangListLogic {
	return &GetLangListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLangListLogic) GetLangList(in *pb.GetLangListReq) (*pb.GetLangListResp, error) {
	langList, err := l.svcCtx.LangModel.FindAll(l.ctx)
	if err != nil {
		return nil, status.New(codes.Canceled, "查询失败").Err()
	}
	list := make([]*pb.LangData, 0)
	for _, v := range langList {
		list = append(list, &pb.LangData{
			Id:       v.Id,
			Lang:     v.Lang,
			Code:     v.Code,
			Original: v.Original != 0,
			Results:  v.Results != 0,
		})
	}
	return &pb.GetLangListResp{
		LangList: list,
	}, nil
}
