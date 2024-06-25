// Code generated by goctl. DO NOT EDIT.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	groupMemberFieldNames          = builder.RawFieldNames(&GroupMember{})
	groupMemberRows                = strings.Join(groupMemberFieldNames, ",")
	groupMemberRowsExpectAutoSet   = strings.Join(stringx.Remove(groupMemberFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	groupMemberRowsWithPlaceHolder = strings.Join(stringx.Remove(groupMemberFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheGroupMemberIdPrefix = "cache:groupMember:id:"
)

type (
	groupMemberModel interface {
		Insert(ctx context.Context, data *GroupMember) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*GroupMember, error)
		Update(ctx context.Context, data *GroupMember) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultGroupMemberModel struct {
		sqlc.CachedConn
		table string
	}

	GroupMember struct {
		Id          uint64         `db:"id"`
		GroupId     string         `db:"group_id"`
		UserId      string         `db:"user_id"`
		RoleLevel   int64          `db:"role_level"`
		JoinTime    sql.NullTime   `db:"join_time"`
		JoinSource  sql.NullInt64  `db:"join_source"`
		InviterUid  sql.NullString `db:"inviter_uid"`
		OperatorUid sql.NullString `db:"operator_uid"`
	}
)

func newGroupMemberModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultGroupMemberModel {
	return &defaultGroupMemberModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`group_member`",
	}
}

func (m *defaultGroupMemberModel) Delete(ctx context.Context, id uint64) error {
	groupMemberIdKey := fmt.Sprintf("%s%v", cacheGroupMemberIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, groupMemberIdKey)
	return err
}

func (m *defaultGroupMemberModel) FindOne(ctx context.Context, id uint64) (*GroupMember, error) {
	groupMemberIdKey := fmt.Sprintf("%s%v", cacheGroupMemberIdPrefix, id)
	var resp GroupMember
	err := m.QueryRowCtx(ctx, &resp, groupMemberIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupMemberRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGroupMemberModel) Insert(ctx context.Context, data *GroupMember) (sql.Result, error) {
	groupMemberIdKey := fmt.Sprintf("%s%v", cacheGroupMemberIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, groupMemberRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.GroupId, data.UserId, data.RoleLevel, data.JoinTime, data.JoinSource, data.InviterUid, data.OperatorUid)
	}, groupMemberIdKey)
	return ret, err
}

func (m *defaultGroupMemberModel) Update(ctx context.Context, data *GroupMember) error {
	groupMemberIdKey := fmt.Sprintf("%s%v", cacheGroupMemberIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, groupMemberRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.GroupId, data.UserId, data.RoleLevel, data.JoinTime, data.JoinSource, data.InviterUid, data.OperatorUid, data.Id)
	}, groupMemberIdKey)
	return err
}

func (m *defaultGroupMemberModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheGroupMemberIdPrefix, primary)
}

func (m *defaultGroupMemberModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupMemberRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultGroupMemberModel) tableName() string {
	return m.table
}
