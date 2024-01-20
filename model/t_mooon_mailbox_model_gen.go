// Code generated by goctl. DO NOT EDIT.

package model

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
	tMooonMailboxFieldNames          = builder.RawFieldNames(&TMooonMailbox{})
	tMooonMailboxRows                = strings.Join(tMooonMailboxFieldNames, ",")
	tMooonMailboxRowsExpectAutoSet   = strings.Join(stringx.Remove(tMooonMailboxFieldNames, "`f_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tMooonMailboxRowsWithPlaceHolder = strings.Join(stringx.Remove(tMooonMailboxFieldNames, "`f_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	tMooonMailboxModel interface {
		Insert(ctx context.Context, data *TMooonMailbox) (sql.Result, error)
		FindOne(ctx context.Context, fId int64) (*TMooonMailbox, error)
		Update(ctx context.Context, data *TMooonMailbox) error
		Delete(ctx context.Context, fId int64) error
	}

	defaultTMooonMailboxModel struct {
		conn  sqlx.SqlConn
		table string
	}

	TMooonMailbox struct {
		FId          int64     `db:"f_id"`           // 自增 ID（信件 ID）
		FRecipient   string    `db:"f_recipient"`    // 收件人
		FDeliverTime time.Time `db:"f_deliver_time"` // 投递时间
		FArrivalTime time.Time `db:"f_arrival_time"` // 到达时间
		FState       int64     `db:"f_state"`        // 信件状态（0 未读，1 已读）
		FLetterBody  string    `db:"f_letter_body"`  // 信件内容
	}
)

func newTMooonMailboxModel(conn sqlx.SqlConn) *defaultTMooonMailboxModel {
	return &defaultTMooonMailboxModel{
		conn:  conn,
		table: "`t_mooon_mailbox`",
	}
}

func (m *defaultTMooonMailboxModel) Delete(ctx context.Context, fId int64) error {
	query := fmt.Sprintf("delete from %s where `f_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, fId)
	return err
}

func (m *defaultTMooonMailboxModel) FindOne(ctx context.Context, fId int64) (*TMooonMailbox, error) {
	query := fmt.Sprintf("select %s from %s where `f_id` = ? limit 1", tMooonMailboxRows, m.table)
	var resp TMooonMailbox
	err := m.conn.QueryRowCtx(ctx, &resp, query, fId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTMooonMailboxModel) Insert(ctx context.Context, data *TMooonMailbox) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, tMooonMailboxRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.FRecipient, data.FDeliverTime, data.FArrivalTime, data.FState, data.FLetterBody)
	return ret, err
}

func (m *defaultTMooonMailboxModel) Update(ctx context.Context, data *TMooonMailbox) error {
	query := fmt.Sprintf("update %s set %s where `f_id` = ?", m.table, tMooonMailboxRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.FRecipient, data.FDeliverTime, data.FArrivalTime, data.FState, data.FLetterBody, data.FId)
	return err
}

func (m *defaultTMooonMailboxModel) tableName() string {
	return m.table
}
