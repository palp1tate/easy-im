package logic

import (
	"context"
	"database/sql"
	"time"

	"github.com/palp1tate/easy-im/apps/social/models"
	"github.com/palp1tate/easy-im/apps/social/rpc/internal/svc"
	"github.com/palp1tate/easy-im/apps/social/rpc/social"
	"github.com/palp1tate/easy-im/pkg/constants"
	"github.com/palp1tate/easy-im/pkg/errorx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInLogic) FriendPutIn(in *social.FriendPutInReq) (*social.FriendPutInResp, error) {
	// 申请人是否与目标是好友关系
	friend, err := l.svcCtx.FriendModel.FindByUidAndFid(l.ctx, in.UserId, in.ReqUid)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return nil, errors.Wrapf(errorx.NewDBErr(), "find friend by uid and fid err %v req %v ", err, in)
	}
	if friend != nil {
		return &social.FriendPutInResp{}, err
	}

	// 是否已经有过申请，申请是不成功，没有完成
	friendReq, err := l.svcCtx.FriendRequestModel.FindByReqUidAndUserId(l.ctx, in.ReqUid, in.UserId)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return nil, errors.Wrapf(errorx.NewDBErr(), "find friendRequest by rid and uid err %v req %v ", err, in)
	}
	if friendReq != nil {
		return &social.FriendPutInResp{}, err
	}

	// 创建申请记录
	_, err = l.svcCtx.FriendRequestModel.Insert(l.ctx, &models.FriendRequest{
		UserId: in.UserId,
		ReqUid: in.ReqUid,
		ReqMsg: sql.NullString{
			Valid:  true,
			String: in.ReqMsg,
		},
		ReqTime: time.Unix(in.ReqTime, 0),
		HandleResult: sql.NullInt64{
			Int64: int64(constants.NoHandlerResult),
			Valid: true,
		},
	})
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDBErr(), "insert friendRequest err %v req %v ", err, in)
	}

	return &social.FriendPutInResp{}, nil
}
