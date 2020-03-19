package middleware

import (
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	if r.Method=="OPTIONS"{
		r.Response.Status=200
		r.ExitAll()
	}
	r.Middleware.Next()
}
