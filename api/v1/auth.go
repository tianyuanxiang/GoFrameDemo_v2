package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

// 登录
type AuthLoginReq struct {
	g.Meta `path:"/login" method:"post"`
}

type AuthLoginRes struct {
	Message string    `json:"message"`
	Token   string    `json:"token"`
	Expire  time.Time `json:"expire"`
}

// 刷新
type AuthRefreshTokenReq struct {
	g.Meta `path:"/refresh_token" method:"post"`
}

type AuthRefreshTokenRes struct {
	Message string    `json:"message"`
	Token   string    `json:"token"`
	Expire  time.Time `json:"expire"`
}

// 登出
type AuthLogoutReq struct {
	g.Meta `path:"/logout" method:"post"`
}

type AuthLogoutRes struct {
	Message string `json:"message"`
}

// 查看用户信息
type UserGetInfoReq struct {
	g.Meta `path:"/user/info" method:"get"`
}

type UserGetInfoRes struct {
	Id          int    `json:"id"`
	IdentityKey string `json:"identity_key"`
	Payload     string `json:"payload"`
}
