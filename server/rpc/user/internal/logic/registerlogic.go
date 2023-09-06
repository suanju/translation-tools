package logic

import (
	"context"

	"translation/rpc/user/internal/svc"
	"translation/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.UserRegisterReq) (*pb.UserRegisterResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserRegisterResp{}, nil
}
