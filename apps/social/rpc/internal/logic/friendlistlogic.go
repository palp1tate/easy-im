package logic

import (
	"context"

	"github.com/palp1tate/easy-im/apps/social/rpc/internal/svc"
	"github.com/palp1tate/easy-im/apps/social/rpc/social"
	"github.com/palp1tate/easy-im/pkg/errorx"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *social.FriendListReq) (*social.FriendListResp, error) {
	friendList, err := l.svcCtx.FriendModel.ListByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDBErr(), "list friend by uid err %v req %v ", err,
			in.UserId)
	}

	var respList []*social.Friend
	err = copier.Copy(&respList, &friendList)
	if err != nil {
		return nil, err
	}

	return &social.FriendListResp{
		List: respList,
	}, nil
}
