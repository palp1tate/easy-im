package friend

import (
	"context"

	"github.com/palp1tate/easy-im/apps/social/api/internal/svc"
	"github.com/palp1tate/easy-im/apps/social/api/internal/types"
	"github.com/palp1tate/easy-im/apps/social/rpc/socialclient"
	"github.com/palp1tate/easy-im/apps/user/rpc/userclient"
	"github.com/palp1tate/easy-im/pkg/jwtx"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewFriendListLogic 好友列表
func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListReq) (resp *types.FriendListResp, err error) {
	uid := jwtx.GetUId(l.ctx)

	friends, err := l.svcCtx.SocialClient.FriendList(l.ctx, &socialclient.FriendListReq{
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}

	if len(friends.List) == 0 {
		return &types.FriendListResp{}, nil
	}

	// 根据好友id获取好友信息
	uidList := make([]string, 0, len(friends.List))
	for _, i := range friends.List {
		uidList = append(uidList, i.FriendUid)
	}

	// 根据uidList查询用户信息
	users, err := l.svcCtx.UserClient.FindUser(l.ctx, &userclient.FindUserReq{
		Ids: uidList,
	})
	if err != nil {
		return &types.FriendListResp{}, nil
	}

	userRecords := make(map[string]*userclient.UserEntity, len(users.User))
	for i := range users.User {
		userRecords[users.User[i].Id] = users.User[i]
	}

	respList := make([]*types.Friend, 0, len(friends.List))
	for _, v := range friends.List {
		friend := &types.Friend{
			Id:        v.Id,
			FriendUid: v.FriendUid,
		}

		if u, ok := userRecords[v.FriendUid]; ok {
			friend.Nickname = u.Nickname
			friend.Avatar = u.Avatar
		}
		respList = append(respList, friend)
	}

	return &types.FriendListResp{
		List: respList,
	}, nil
}
