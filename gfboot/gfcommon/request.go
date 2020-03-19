package gfcommon

import (
	"gf-app/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

func requestData(r *ghttp.Request, params map[string]interface{}, method string) map[string]interface{} {
	fields := g.Map{}
	valicateRules := g.Map{}
	for k, v := range params {
		fields[k] = ""
		if v != "" {
			valicateRules[k] = v
		}
	}
	data := g.Map{}
	if method == "get" {
		data = r.GetQueryMap(fields)
	} else {
		data = r.GetPostMap(fields)
	}
	if e := gvalid.CheckMap(data, valicateRules); e != nil {
		response.JsonFailGError(r, e, nil)
	}
	return data
}

func RequestGetData(r *ghttp.Request, params map[string]interface{}) map[string]interface{} {
	return requestData(r, params, "get")
}

func RequestPostData(r *ghttp.Request, params map[string]interface{}) map[string]interface{} {
	return requestData(r, params, "post")
}
