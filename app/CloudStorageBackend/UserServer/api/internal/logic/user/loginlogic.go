package user

import (
	"UserServer/api/internal/svc"
	"UserServer/api/internal/types"
	"UserServer/api/units"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.CssUserModel.FindOneByAccount(l.ctx, req.Name, units.Md5(req.Password))
	token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, user.Id, user.Account, user.Name)

	return &types.LoginRes{
		AccessToken: token,
		Result:      "true",
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, id int64, account string, name string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + 60000
	claims["id"] = id
	claims["account"] = account
	claims["name"] = name
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
