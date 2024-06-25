package logic

import (
	"context"

	"github.com/palp1tate/easy-im/apps/social/models"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/palp1tate/easy-im/apps/social/rpc/internal/svc"
	"github.com/palp1tate/easy-im/apps/social/rpc/social"
	"github.com/palp1tate/easy-im/pkg/constants"
	"github.com/palp1tate/easy-im/pkg/errorx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrFriendReqBeforePass   = errorx.NewMsg("好友申请并已经通过")
	ErrFriendReqBeforeRefuse = errorx.NewMsg("好友申请已经被拒绝")
)

type FriendPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInHandleLogic) FriendPutInHandle(in *social.FriendPutInHandleReq) (*social.FriendPutInHandleResp, error) {
	// 获取好友申请记录
	friendReq, err := l.svcCtx.FriendRequestModel.FindOne(l.ctx, uint64(in.FriendReqId))
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDBErr(), "find friendRequest by friendReqid err %v req %v ", err,
			in.FriendReqId)
	}

	// 验证是否有处理
	switch constants.HandlerResult(friendReq.HandleResult.Int64) {
	case constants.PassHandlerResult:
		return nil, errors.WithStack(ErrFriendReqBeforePass)
	case constants.RefuseHandlerResult:
		return nil, errors.WithStack(ErrFriendReqBeforeRefuse)
	default:
		break
	}

	// 修改申请结果 -》 通过【建立两条好友关系记录】 -》 事务
	err = l.svcCtx.FriendRequestModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		friendReq.HandleResult.Int64 = int64(in.HandleResult)
		if err := l.svcCtx.FriendRequestModel.Update(l.ctx, session, friendReq); err != nil {
			return errors.Wrapf(errorx.NewDBErr(), "update friend request err %v, req %v", err, friendReq)
		}

		if constants.HandlerResult(in.HandleResult) != constants.PassHandlerResult {
			return nil
		}
		// 先判断两人是否已经是好友
		friend, err := l.svcCtx.FriendModel.FindByUidAndFidWithSession(l.ctx, session, friendReq.UserId, friendReq.ReqUid)
		if err != nil && !errors.Is(err, models.ErrNotFound) {
			return errors.Wrapf(errorx.NewDBErr(), "find friend by uid and fid err %v req %v ", err, friendReq)
		}
		if friend == nil {
			friends := []*models.Friend{
				{
					UserId:    friendReq.UserId,
					FriendUid: friendReq.ReqUid,
				}, {
					UserId:    friendReq.ReqUid,
					FriendUid: friendReq.UserId,
				},
			}

			_, err = l.svcCtx.FriendModel.Inserts(l.ctx, session, friends...)
			if err != nil {
				return errors.Wrapf(errorx.NewDBErr(), "friends inserts err %v, req %v", err, friends)
			}
		}
		return nil
	})

	return &social.FriendPutInHandleResp{}, err
}
