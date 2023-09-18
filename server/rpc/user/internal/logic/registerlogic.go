package logic

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"github.com/bwmarrin/snowflake"
	"translation/rpc/user/model/users"

	"translation/rpc/user/internal/svc"
	"translation/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	//验证邮箱
	if isRegister, _ := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email); isRegister != nil {
		return nil, status.New(codes.Canceled, "邮箱已注册").Err()
	}
	node, _ := snowflake.NewNode(1)
	id := node.Generate().Int64()
	salt := node.Generate().Base32()
	hash := sha256.New()
	hash.Write([]byte(salt + in.Password + salt))
	password := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	_, err := l.svcCtx.UserModel.Insert(l.ctx, &users.TnUsers{
		Id:           id,
		Email:        in.Email,
		Password:     password,
		PasswordSalt: salt,
		Status:       0,
	})
	if err != nil {
		return nil, status.New(codes.Canceled, "注册失败").Err()
	}
	return &pb.UserRegisterResp{
		UserId: id,
	}, nil
}
