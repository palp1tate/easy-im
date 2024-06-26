package friend

import (
	"context"

	"github.com/palp1tate/easy-im/apps/social/api/internal/svc"
	"github.com/palp1tate/easy-im/apps/social/api/internal/types"
	"github.com/palp1tate/easy-im/apps/social/rpc/socialclient"
	"github.com/palp1tate/easy-im/pkg/jwtx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewFriendPutInListLogic 好友申请列表
func NewFriendPutInListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInListLogic {
	return &FriendPutInListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendPutInListLogic) FriendPutInList(req *types.FriendPutInListReq) (resp *types.FriendPutInListResp, err error) {
	list, err := l.svcCtx.SocialClient.FriendPutInList(l.ctx, &socialclient.FriendPutInListReq{
		UserId: jwtx.GetUId(l.ctx),
	})
	if err != nil {
		return nil, err
	}

	var respList []*types.FriendRequest
	err = copier.Copy(&respList, list.List)
	if err != nil {
		return nil, err
	}

	return &types.FriendPutInListResp{List: respList}, nil
}
