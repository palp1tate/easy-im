package logic

import (
	"context"
	"time"

	"github.com/palp1tate/easy-im/apps/user/models"
	"github.com/palp1tate/easy-im/apps/user/rpc/internal/svc"
	"github.com/palp1tate/easy-im/apps/user/rpc/user"
	"github.com/palp1tate/easy-im/pkg/encrypt"
	"github.com/palp1tate/easy-im/pkg/errorx"
	"github.com/palp1tate/easy-im/pkg/jwtx"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrPhoneNotRegister = errorx.New(errorx.ServerCommonError, "手机号码未注册")
	ErrUserPwdError     = errorx.New(errorx.ServerCommonError, "密码错误")
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// 验证用户是否注册，根据手机号码验证
	userEntity, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, errors.WithStack(ErrPhoneNotRegister)
		}
		return nil, errors.Wrapf(errorx.NewDBErr(), "find user by phone err %v , req %v", err, in.Phone)
	}

	// 密码验证
	if !encrypt.ValidatePasswordHash(in.Password, userEntity.Password.String) {
		return nil, errors.WithStack(ErrUserPwdError)
	}

	// 生成token
	now := time.Now().Unix()
	token, err := jwtx.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire,
		userEntity.Id)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewInternalErr(), "jwtx get jwt token err %v", err)
	}

	return &user.LoginResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
