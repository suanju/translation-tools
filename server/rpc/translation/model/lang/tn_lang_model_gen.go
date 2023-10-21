// Code generated by goctl. DO NOT EDIT.

package lang

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tnLangFieldNames          = builder.RawFieldNames(&TnLang{})
	tnLangRows                = strings.Join(tnLangFieldNames, ",")
	tnLangRowsExpectAutoSet   = strings.Join(stringx.Remove(tnLangFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tnLangRowsWithPlaceHolder = strings.Join(stringx.Remove(tnLangFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tnLangModel interface {
		Insert(ctx context.Context, data *TnLang) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TnLang, error)
		Update(ctx context.Context, data *TnLang) error
		Delete(ctx context.Context, id int64) error
		FindAll(ctx context.Context) ([]*TnLang, error)
	}

	defaultTnLangModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TnLang struct {
		Id       int64        `db:"id"`
		Lang     string       `db:"lang"`      // 语言名称
		Code     string       `db:"code"`      // 语言代码
		Original int64        `db:"original"`  // 是否可作用原始语言  0 ->否 1-> 是
		Results  int64        `db:"results"`   // 是否可作用结果语言  0 ->否 1-> 是
		Status   int64        `db:"status"`    //  0->禁用  1->可用
		CreateAt time.Time    `db:"create_at"` // 创建时间
		UpdateAt sql.NullTime `db:"update_at"` // 更新时间
	}
)

func newTnLangModel(conn sqlx.SqlConn) *defaultTnLangModel {
	return &defaultTnLangModel{
		conn:  conn,
		table: "`tn_lang`",
	}
}

func (m *defaultTnLangModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultTnLangModel) FindOne(ctx context.Context, id int64) (*TnLang, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tnLangRows, m.table)
	var resp TnLang
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTnLangModel) Insert(ctx context.Context, data *TnLang) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, tnLangRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Lang, data.Code, data.Original, data.Results, data.Status)
	return ret, err
}

func (m *defaultTnLangModel) Update(ctx context.Context, data *TnLang) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tnLangRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Lang, data.Code, data.Original, data.Results, data.Status, data.Id)
	return err
}

func (m *defaultTnLangModel) FindAll(ctx context.Context) ([]*TnLang , error) {
	query := fmt.Sprintf("select %s from %s", tnLangRows, m.table)
	var resp []*TnLang
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTnLangModel) tableName() string {
	return m.table
}