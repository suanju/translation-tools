package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"translation/rpc/user/internal/config"
	"translation/rpc/user/model/users"
)

type ServiceContext struct {
	Config    config.Config
	UserModel users.TnUsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:    c,
		UserModel: users.NewTnUsersModel(sqlConn),
	}
}
