package cmd

import (
	"context"
	"demo/internal/service"

	"demo/internal/controller"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/api", func(group *ghttp.RouterGroup) {
				// 该中间件应用于全局接口
				group.Middleware(
					service.Middle().MiddlewareCORS,  // CORS权限
					service.Middle().ResponseHandler, // 返回处理
				)
				// 第一个子组
				// 该分组没有使用中间件
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.Bind(
						// 该组没有编写特定的中间件，意味着不用判断鉴权
						// New 结构体绑定了很多方法，可以统一注册
						controller.New(),               // 图书管理
						controller.TypeCtl,             // 图书类型管理
						controller.UserCtl,             // 用户管理
						controller.UBorrow,             // 用户查询
						controller.BBInformation,       // 用户借阅信息查询
						controller.UHistory,            // 用户历史借阅信息查询
						controller.UBooksBorrowed,      // 用户当前借阅信息查询
						controller.UPersonalProfile,    // 用户个人信息
						controller.ReaderManagementCtl, //读者信息管理
					)
				})
				group.Group("/", func(group *ghttp.RouterGroup) {
					// Auth的鉴权登录
					group.Middleware(
						//service.Middle().Auth,
						ghttp.MiddlewareHandlerResponse,
					)
					group.Bind(
						controller.Login, //用户登录
					)
					group.Group("/", func(group *ghttp.RouterGroup) {
						// Auth的鉴权登录
						group.Middleware(service.Middle().Auth)
						group.ALLMap(g.Map{
							"/user/info": controller.User.Info,
						})
					})

				})
				//group.Group("/v1/user/info", func(group *ghttp.RouterGroup) {
				//	// Auth的鉴权登录
				//	group.Middleware(
				//		service.Middle().Auth,
				//	)
				//	group.Bind(
				//		controller.Login, //用户登录
				//	)
				//})
			})
			s.Run()
			return nil
		},
	}
)
