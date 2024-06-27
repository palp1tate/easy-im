package logic

import (
	"context"
	"database/sql"
	"time"

	"github.com/palp1tate/easy-im/apps/user/models"
	"github.com/palp1tate/easy-im/apps/user/rpc/internal/svc"
	"github.com/palp1tate/easy-im/apps/user/rpc/user"
	"github.com/palp1tate/easy-im/pkg/encrypt"
	"github.com/palp1tate/easy-im/pkg/errorx"
	"github.com/palp1tate/easy-im/pkg/jwtx"
	"github.com/palp1tate/easy-im/pkg/wuid"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrPhoneIsRegister = errorx.New(errorx.ServerCommonError, "手机号码已注册")

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	userEntity, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, in.Phone)
	if err != nil && !errors.Is(err, models.ErrNotFound) {
		return nil, errors.Wrapf(errorx.NewDBErr(), "find user by phone err %v , req %v", err, in.Phone)
	}
	if userEntity != nil {
		return nil, errors.WithStack(ErrPhoneIsRegister)
	}
	userEntity = &models.User{
		Id:       wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
		Avatar:   in.Avatar,
		Nickname: in.Nickname,
		Phone:    in.Phone,
		Sex:      sql.NullInt64{Int64: int64(in.Sex), Valid: true},
	}
	if len(in.Password) > 0 {
		genPassword, err := encrypt.GenPasswordHash([]byte(in.Password))
		if err != nil {
			return nil, errors.Wrapf(errorx.NewInternalErr(), "gen password hash err %v", err)
		}
		userEntity.Password = sql.NullString{
			String: string(genPassword),
			Valid:  true,
		}
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, userEntity)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewDBErr(), "insert user err %v , req %v", err, in)
	}

	now := time.Now().Unix()
	token, err := jwtx.GetJwtToken(l.svcCtx.Config.Jwt.AccessSecret, now, l.svcCtx.Config.Jwt.AccessExpire, userEntity.Id)
	if err != nil {
		return nil, errors.Wrapf(errorx.NewInternalErr(), "jwtx get jwt token err %v", err)
	}
	return &user.RegisterResp{
		Token:  token,
		Expire: now + l.svcCtx.Config.Jwt.AccessExpire,
	}, nil
}
