package middleware

import (
	"gf-app/library/auth"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

var whiteUrls *garray.StrArray

func init() {
	whiteUrls = garray.NewStrArrayFrom(g.Config().GetStrings("jwt.whiteUrls"))
}

func MiddlewareJWT(r *ghttp.Request) {
	if r.Method=="OPTIONS"{
		r.Response.Status=http.StatusNoContent
		r.ExitAll()
	}
	if !whiteUrls.Contains(r.URL.Path) {
		auth.GfJWTMiddleware.MiddlewareFunc()(r)
	}
	r.Middleware.Next()
}
