package friend

import (
	"context"

	"github.com/palp1tate/easy-im/apps/social/api/internal/svc"
	"github.com/palp1tate/easy-im/apps/social/api/internal/types"
	"github.com/palp1tate/easy-im/apps/social/rpc/socialclient"
	"github.com/palp1tate/easy-im/pkg/jwtx"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewFriendPutInHandleLogic 好友申请处理
func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendPutInHandleLogic) FriendPutInHandle(req *types.FriendPutInHandleReq) (resp *types.FriendPutInHandleResp, err error) {
	_, err = l.svcCtx.SocialClient.FriendPutInHandle(l.ctx, &socialclient.FriendPutInHandleReq{
		FriendReqId:  req.FriendReqId,
		UserId:       jwtx.GetUId(l.ctx),
		HandleResult: req.HandleResult,
	})

	return
}
