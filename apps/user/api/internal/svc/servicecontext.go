package svc

import (
	"github.com/palp1tate/easy-im/apps/user/api/internal/config"
	"github.com/palp1tate/easy-im/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserClient userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserClient: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
