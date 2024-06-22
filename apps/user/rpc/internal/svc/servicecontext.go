package svc

import (
	"github.com/palp1tate/easy-im/apps/user/models"
	"github.com/palp1tate/easy-im/apps/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	models.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: models.NewUserModel(sqlConn, c.Cache),
	}
}
