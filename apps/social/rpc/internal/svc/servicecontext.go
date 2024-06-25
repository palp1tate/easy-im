package svc

import (
	"github.com/palp1tate/easy-im/apps/social/models"
	"github.com/palp1tate/easy-im/apps/social/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	models.FriendModel
	models.FriendRequestModel
	models.GroupModel
	models.GroupRequestModel
	models.GroupMemberModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:             c,
		FriendModel:        models.NewFriendModel(sqlConn, c.Cache),
		FriendRequestModel: models.NewFriendRequestModel(sqlConn, c.Cache),
		GroupModel:         models.NewGroupModel(sqlConn, c.Cache),
		GroupRequestModel:  models.NewGroupRequestModel(sqlConn, c.Cache),
		GroupMemberModel:   models.NewGroupMemberModel(sqlConn, c.Cache),
	}
}
