package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"translation/rpc/translation/internal/config"
	"translation/rpc/translation/model/lang"
)

type ServiceContext struct {
	Config    config.Config
	LangModel lang.TnLangModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config:    c,
		LangModel: lang.NewTnLangModel(sqlConn),
	}
}
