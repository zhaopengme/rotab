package router

import (
	_ "gf-app/app/module/admin/controller"
	_ "gf-app/app/module/front/tpl"
	"gf-app/library/auth"
	"gf-app/router/middleware"
	"github.com/gogf/gf/frame/g"
)

// 统一路由注册.
func init() {
	s := g.Server()
	s.BindMiddleware("/file/upload", middleware.MiddlewareCORS)
	s.BindMiddleware("/admin/*/**", middleware.MiddlewareCORS)
	s.BindMiddleware("/admin/*/**", middleware.MiddlewareJWT)
	s.BindHandler("/admin/login", auth.GfJWTMiddleware.LoginHandler)
	s.BindHandler("/admin/refresh_token", auth.GfJWTMiddleware.RefreshHandler)
}
