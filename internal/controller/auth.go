package controller

import (
	"context"
	v1 "demo/api/v1"
	"demo/internal/logic/auth"
	"fmt"
)

type loginController struct{}

var Login = loginController{}

func (c *loginController) AuthLogin(ctx context.Context, req *v1.AuthLoginReq) (res *v1.AuthLoginRes, err error) {
	fmt.Println("sssssss")
	res = &v1.AuthLoginRes{}
	res.Message = "登录信息如下:"
	res.Token, res.Expire = auth.Auth().LoginHandler(ctx)
	return
}

func (c *loginController) RefreshToken(ctx context.Context, req *v1.AuthRefreshTokenReq) (res *v1.AuthRefreshTokenRes, err error) {
	res = &v1.AuthRefreshTokenRes{}
	res.Token, res.Expire = auth.Auth().RefreshHandler(ctx)
	return
}

func (c *loginController) AuthLogout(ctx context.Context, req *v1.AuthLogoutReq) (res *v1.AuthLogoutRes, err error) {
	fmt.Println("Logoutt")
	res = &v1.AuthLogoutRes{
		Message: "用户已登出",
	}
	auth.Auth().LogoutHandler(ctx)
	return
}
