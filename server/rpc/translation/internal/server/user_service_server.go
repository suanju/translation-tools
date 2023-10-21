// Code generated by goctl. DO NOT EDIT.
// Source: translation.proto

package server

import (
	"context"

	"translation/rpc/translation/internal/logic"
	"translation/rpc/translation/internal/svc"
	"translation/rpc/translation/pb"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) GetLangList(ctx context.Context, in *pb.GetLangListReq) (*pb.GetLangListResp, error) {
	l := logic.NewGetLangListLogic(ctx, s.svcCtx)
	return l.GetLangList(in)
}