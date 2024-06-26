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
	friendFieldNames          = builder.RawFieldNames(&Friend{})
	friendRows                = strings.Join(friendFieldNames, ",")
	friendRowsExpectAutoSet   = strings.Join(stringx.Remove(friendFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	friendRowsWithPlaceHolder = strings.Join(stringx.Remove(friendFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheFriendIdPrefix = "cache:friend:id:"
)

type (
	friendModel interface {
		Insert(ctx context.Context, data *Friend) (sql.Result, error)
		Inserts(ctx context.Context, session sqlx.Session, data ...*Friend) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*Friend, error)
		FindByUidAndFid(ctx context.Context, uid, fid string) (*Friend, error)
		FindByUidAndFidWithSession(ctx context.Context, session sqlx.Session, uid, fid string) (*Friend, error)
		ListByUserId(ctx context.Context, userId string) ([]*Friend, error)
		Update(ctx context.Context, data *Friend) error
		Delete(ctx context.Context, id uint64) error
	}

	defaultFriendModel struct {
		sqlc.CachedConn
		table string
	}

	Friend struct {
		Id        uint64         `db:"id"`
		UserId    string         `db:"user_id"`
		FriendUid string         `db:"friend_uid"`
		Remark    sql.NullString `db:"remark"`
		AddSource sql.NullInt64  `db:"add_source"`
		CreatedAt sql.NullTime   `db:"created_at"`
	}
)

func newFriendModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultFriendModel {
	return &defaultFriendModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`friend`",
	}
}

func (m *defaultFriendModel) Delete(ctx context.Context, id uint64) error {
	friendIdKey := fmt.Sprintf("%s%v", cacheFriendIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, friendIdKey)
	return err
}

func (m *defaultFriendModel) FindOne(ctx context.Context, id uint64) (*Friend, error) {
	friendIdKey := fmt.Sprintf("%s%v", cacheFriendIdPrefix, id)
	var resp Friend
	err := m.QueryRowCtx(ctx, &resp, friendIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", friendRows, m.table)
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

func (m *defaultFriendModel) FindByUidAndFid(ctx context.Context, uid, fid string) (*Friend, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `friend_uid` = ?", friendRows, m.table)

	var resp Friend
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, uid, fid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFriendModel) FindByUidAndFidWithSession(ctx context.Context, session sqlx.Session, uid, fid string) (*Friend, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `friend_uid` = ?", friendRows, m.table)

	var resp Friend
	err := session.QueryRowCtx(ctx, &resp, query, uid, fid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFriendModel) ListByUserId(ctx context.Context, userId string) ([]*Friend, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? ", friendRows, m.table)

	var resp []*Friend
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultFriendModel) Insert(ctx context.Context, data *Friend) (sql.Result, error) {
	friendIdKey := fmt.Sprintf("%s%v", cacheFriendIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, friendRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.FriendUid, data.Remark, data.AddSource)
	}, friendIdKey)
	return ret, err
}

func (m *defaultFriendModel) Inserts(ctx context.Context, session sqlx.Session, data ...*Friend) (sql.Result, error) {
	var (
		sql  strings.Builder
		args []any
	)

	if len(data) == 0 {
		return nil, nil
	}

	// insert into tablename values(数据), (数据)
	sql.WriteString(fmt.Sprintf("insert into %s (%s) values ", m.table, friendRowsExpectAutoSet))

	for i, v := range data {
		sql.WriteString("(?, ?, ?, ?)")
		args = append(args, v.UserId, v.FriendUid, v.Remark.String, v.AddSource.Int64)
		if i == len(data)-1 {
			break
		}
		sql.WriteString(",")
	}
	return session.ExecCtx(ctx, sql.String(), args...)
}

func (m *defaultFriendModel) Update(ctx context.Context, data *Friend) error {
	friendIdKey := fmt.Sprintf("%s%v", cacheFriendIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, friendRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.FriendUid, data.Remark, data.AddSource, data.Id)
	}, friendIdKey)
	return err
}

func (m *defaultFriendModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheFriendIdPrefix, primary)
}

func (m *defaultFriendModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", friendRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFriendModel) tableName() string {
	return m.table
}