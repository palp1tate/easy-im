package logic

import (
	"context"

	"github.com/palp1tate/easy-im/apps/social/models"

	"github.com/palp1tate/easy-im/pkg/errorx"
	"github.com/pkg/errors"

	"github.com/palp1tate/easy-im/apps/social/rpc/internal/svc"
	"github.com/palp1tate/easy-im/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInListLogic {
	return &FriendPutInListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInListLogic) FriendPutInList(in *social.FriendPutInListReq) (*social.FriendPutInListResp, error) {
	friendReqList, err := l.svcCtx.FriendRequestModel.ListFriendRequest(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDBErr(), "find list friend req err %v req %v", err, in.UserId)
	}

	resp, _ := convertFriendRequests(friendReqList)
	return &social.FriendPutInListResp{
		List: resp,
	}, nil
}

func convertFriendRequests(dbRequests []*models.FriendRequest) ([]*social.FriendRequest, error) {
	res := make([]*social.FriendRequest, len(dbRequests))
	for i, dbReq := range dbRequests {
		res[i] = &social.FriendRequest{
			Id:           int32(dbReq.Id),
			UserId:       dbReq.UserId,
			ReqUid:       dbReq.ReqUid,
			ReqMsg:       dbReq.ReqMsg.String,
			ReqTime:      dbReq.ReqTime.Unix(),
			HandleResult: int32(dbReq.HandleResult.Int64),
		}
	}
	return res, nil
}
