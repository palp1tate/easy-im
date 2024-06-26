package svc

import (
	"github.com/palp1tate/easy-im/apps/social/api/internal/config"
	"github.com/palp1tate/easy-im/apps/social/rpc/socialclient"
	"github.com/palp1tate/easy-im/apps/user/rpc/userclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config       config.Config
	SocialClient socialclient.Social
	UserClient   userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		SocialClient: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
		UserClient:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
