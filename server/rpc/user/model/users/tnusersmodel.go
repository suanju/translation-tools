package users

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TnUsersModel = (*customTnUsersModel)(nil)

type (
	// TnUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTnUsersModel.
	TnUsersModel interface {
		tnUsersModel
	}

	customTnUsersModel struct {
		*defaultTnUsersModel
	}
)

// NewTnUsersModel returns a model for the database table.
func NewTnUsersModel(conn sqlx.SqlConn) TnUsersModel {
	return &customTnUsersModel{
		defaultTnUsersModel: newTnUsersModel(conn),
	}
}
