package friend

import (
	"context"

	"github.com/palp1tate/easy-im/apps/social/api/internal/svc"
	"github.com/palp1tate/easy-im/apps/social/api/internal/types"
	"github.com/palp1tate/easy-im/apps/social/rpc/socialclient"
	"github.com/palp1tate/easy-im/pkg/jwtx"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewFriendPutInLogic 好友申请
func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendPutInLogic) FriendPutIn(req *types.FriendPutInReq) (resp *types.FriendPutInResp, err error) {
	uid := jwtx.GetUId(l.ctx)

	_, err = l.svcCtx.SocialClient.FriendPutIn(l.ctx, &socialclient.FriendPutInReq{
		UserId:  uid,
		ReqUid:  req.UserId,
		ReqMsg:  req.ReqMsg,
		ReqTime: req.ReqTime,
	})

	return
}
