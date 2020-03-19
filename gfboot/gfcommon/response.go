package gfcommon

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

func JsonFailGError(r *ghttp.Request, e *gvalid.Error, data interface{}) {
	if e != nil {
		json(r, -1, e.String(), data)
	}
}

func JsonFailError(r *ghttp.Request, e error, data interface{}) {
	if e != nil {
		json(r, -1, e.Error(), data)
	}
}

func JsonFail(r *ghttp.Request, msg string, data interface{}) {
	json(r, -1, msg, data)
}

func JsonOK(r *ghttp.Request, msg string, data interface{}) {
	json(r, 0, msg, data)
}

func json(r *ghttp.Request, code int, msg string, data interface{}) {
	r.Response.WriteJson(g.Map{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	r.ExitAll()
}
