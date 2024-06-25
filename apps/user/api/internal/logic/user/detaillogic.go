package user

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/palp1tate/easy-im/apps/user/rpc/user"
	"github.com/palp1tate/easy-im/pkg/jwtx"

	"github.com/palp1tate/easy-im/apps/user/api/internal/svc"
	"github.com/palp1tate/easy-im/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// NewDetailLogic 获取用户信息
func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	uid := jwtx.GetUId(l.ctx)

	userInfoResp, err := l.svcCtx.UserClient.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}

	var res types.User
	err = copier.Copy(&res, userInfoResp.User)
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResp{Info: res}, nil
}
