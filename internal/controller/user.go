package controller

import (
	"context"
	"demo/internal/logic/auth"
	"fmt"
	"github.com/gogf/gf-jwt/v2/example/api"
	"github.com/gogf/gf/v2/util/gconv"
)

type userController struct{}

var User = userController{}

// Info should be authenticated to view.
// It is the get user data handler
func (c *userController) Info(ctx context.Context, req *api.UserGetInfoReq) (res *api.UserGetInfoRes, err error) {
	fmt.Println("azazazaz")
	return &api.UserGetInfoRes{
		Id:          gconv.Int(auth.Auth().GetIdentity(ctx)),
		IdentityKey: auth.Auth().IdentityKey,
		Payload:     auth.Auth().GetPayload(ctx),
	}, nil
}
