package logic

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"translation/rpc/user/internal/svc"
	"translation/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.UserLoginReq) (*pb.UserLoginResp, error) {
	//查询用户
	user, _ := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)
	if user == nil {
		return nil, status.New(codes.Canceled, "用户不存在").Err()
	}
	hash := sha256.New()
	hash.Write([]byte(user.PasswordSalt + in.Password + user.PasswordSalt))
	passwordCheck := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	if user.Password != passwordCheck {
		return nil, status.New(codes.Canceled, "密码错误").Err()
	}
	return &pb.UserLoginResp{
		UserId: user.Id,
	}, nil
}
