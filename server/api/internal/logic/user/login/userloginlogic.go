package login

import (
	"context"
	"google.golang.org/grpc/status"
	"strconv"
	"translation/api/internal/utils/errorx"
	"translation/rpc/user/pb"

	"translation/api/internal/svc"
	"translation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	result, err := l.svcCtx.UserService.Login(l.ctx, &pb.UserLoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, errorx.NewDefaultErrorMessage(status.Convert(err).Message())
	}
	//生成token jwt
	token, err := l.svcCtx.AuthService.GetJwtToken(strconv.FormatInt(result.UserId, 10))
	if err != nil {
		return nil, errorx.NewDefaultErrorMessage("登入失败")
	}
	return &types.UserLoginResp{
		Id:    result.UserId,
		Email: req.Email,
		Token: token,
	}, nil
}
