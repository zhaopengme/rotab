package controller

import (
	"gf-app/app/dao"
	"gf-app/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type OptionCtr struct{}

func init() {
	g.Server().BindObject("/admin/option", new(OptionCtr))
}

func (c *OptionCtr) Index(r *ghttp.Request) {
	r.Response.Write("object index")
}

func (c *OptionCtr) GetByCategory(r *ghttp.Request) {
	category := r.GetString("category", "")
	if category == "" {
		response.JsonFail(r, "param error!", "")
	}
	result, e := dao.OptionDao().FindByCategory(category)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", result)
}

func (c *OptionCtr) UpdateGeneralOption(r *ghttp.Request) {
	data := r.GetPostMapStrStr(g.Map{"title": "", "subTitle": "", "description": ""})
	validationRules := map[string]string{
		"title":       "required",
		"subTitle":    "required",
		"description": "required",
	}
	if e := gvalid.CheckMap(data, validationRules); e != nil {
		response.JsonFailGError(r, e, nil)
	}

	for k, value := range data {
		dao.OptionDao().SaveOrUpdate(&g.Map{"category": "general", "name": k, "value": value})
	}

	response.JsonOK(r, "", nil)
}
