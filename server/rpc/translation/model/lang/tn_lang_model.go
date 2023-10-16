package lang

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TnLangModel = (*customTnLangModel)(nil)

type (
	// TnLangModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTnLangModel.
	TnLangModel interface {
		tnLangModel
	}

	customTnLangModel struct {
		*defaultTnLangModel
	}
)

// NewTnLangModel returns a model for the database table.
func NewTnLangModel(conn sqlx.SqlConn) TnLangModel {
	return &customTnLangModel{
		defaultTnLangModel: newTnLangModel(conn),
	}
}
