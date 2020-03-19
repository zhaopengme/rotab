package controller

import (
	"gf-app/app/dao"
	"gf-app/gfboot/gfcommon"
	"gf-app/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

type NavigationsCtr struct{}

func init() {
	g.Server().BindObject("/admin/navigations", new(NavigationsCtr))
}

func (o *NavigationsCtr) Get(r *ghttp.Request) {
	data := gfcommon.RequestGetData(r, g.Map{"id": "required"})
	record, e := dao.NavigationsDao().Find(data)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", record)
}

func (o *NavigationsCtr) List(r *ghttp.Request) {
	page := &gfcommon.Page{
		Limit: r.GetInt("limit", 10),
		Page:  r.GetInt("page", 1),
	}
	result, e := dao.NavigationsDao().Page(page, nil)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", result)
}

func (o *NavigationsCtr) All(r *ghttp.Request) {
	records, e := dao.NavigationsDao().FindAll(nil)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", records)
}

func (o *NavigationsCtr) Delete(r *ghttp.Request) {
	data := gfcommon.RequestGetData(r, g.Map{"id": "required"})
	record, e := dao.NavigationsDao().Delete(data)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", record)
}

func (o *NavigationsCtr) SaveOrUpdate(r *ghttp.Request) {
	data := gfcommon.RequestPostData(r, g.Map{"title": "required", "url": "required", "open_method": "required", "status": "required"})
	id := r.GetPostInt("id", 0)
	data["number"] = r.GetPostInt("number", 0)
	if id == 0 {
		data["created_at"] = gtime.Now()
		_, e := dao.NavigationsDao().Save(data)
		response.JsonFailError(r, e, nil)
	} else {
		data["updated_at"] = gtime.Now()
		_, e := dao.NavigationsDao().Update(data, g.Map{"id": id})
		response.JsonFailError(r, e, nil)
	}
	response.JsonOK(r, "", nil)
}
