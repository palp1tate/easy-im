package logic

import (
	"context"

	"github.com/palp1tate/easy-im/apps/user/models"
	"github.com/palp1tate/easy-im/apps/user/rpc/internal/svc"
	"github.com/palp1tate/easy-im/apps/user/rpc/user"
	"github.com/palp1tate/easy-im/pkg/errorx"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrUserNotFound = errorx.New(errorx.ServerCommonError, "用户不存在")

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	userEntity, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, errors.WithStack(ErrUserNotFound)
		}
		return nil, errors.Wrapf(errorx.NewDBErr(), "find user by id err %v , req %v", err, in.Id)
	}
	var resp user.UserEntity
	err = copier.Copy(&resp, userEntity)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewInternalErr(), "copier copy err %v", err)
	}

	return &user.GetUserInfoResp{
		User: &resp,
	}, nil
}
