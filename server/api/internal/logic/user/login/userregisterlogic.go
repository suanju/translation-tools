package login

import (
	"context"
	"strconv"
	"translation/api/internal/utils/errorx"
	"translation/rpc/user/pb"

	"google.golang.org/grpc/status"

	"translation/api/internal/svc"
	"translation/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterResp, err error) {
	if req.Password != req.PasswordCheck {
		return nil, errorx.NewDefaultErrorMessage("两次密码不一致")
	}
	result, err := l.svcCtx.UserService.Register(l.ctx, &pb.UserRegisterReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, errorx.NewDefaultErrorMessage(status.Convert(err).Message())
	}
	//生成token jwt
	token, err := l.svcCtx.AuthService.GetJwtToken(strconv.FormatInt(result.UserId, 10))
	if err != nil {
		return nil, errorx.NewDefaultErrorMessage("注册失败")
	}
	return &types.UserRegisterResp{
		Id:    result.UserId,
		Email: req.Email,
		Token: token,
	}, nil
}
