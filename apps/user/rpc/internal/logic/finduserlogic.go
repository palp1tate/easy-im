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

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	var (
		users []*models.User
		err   error
	)

	if in.Phone != "" {
		userEntity, err := l.svcCtx.UserModel.FindByPhone(l.ctx, in.Phone)
		if err == nil {
			users = append(users, userEntity)
		}
	} else if in.Name != "" {
		users, err = l.svcCtx.UserModel.ListByName(l.ctx, in.Name)
	} else if len(in.Ids) > 0 {
		users, err = l.svcCtx.UserModel.ListByIds(l.ctx, in.Ids)
	}

	if err != nil {
		return nil, errors.Wrapf(errorx.NewInternalErr(), "find user err %v", err)
	}

	var resp []*user.UserEntity
	err = copier.Copy(&resp, &users)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewInternalErr(), "copier copy err %v", err)
	}

	return &user.FindUserResp{
		User: resp,
	}, nil
}
