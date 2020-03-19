package controller

import (
	"gf-app/app/dao"
	"gf-app/gfboot/gfcommon"
	"gf-app/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type CategoriesCtr struct{}

func init() {
	g.Server().BindObject("/admin/categories", new(CategoriesCtr))
}

func (o *CategoriesCtr) Get(r *ghttp.Request) {
	data := gfcommon.RequestGetData(r, g.Map{"id": "required"})
	record, e := dao.CategoriesDao().Find(data)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", record)
}

func (o *CategoriesCtr) List(r *ghttp.Request) {
	page := &gfcommon.Page{
		Limit:   r.GetInt("limit", 10),
		Page: r.GetInt("page", 1),
	}
	result, e := dao.CategoriesDao().Page(page, nil)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", result)
}

func (o *CategoriesCtr) All(r *ghttp.Request) {
	records, e := dao.CategoriesDao().FindAll(nil)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", records)
}

func (o *CategoriesCtr) Delete(r *ghttp.Request) {
	data := gfcommon.RequestGetData(r, g.Map{"id": "required"})
	record, e := dao.CategoriesDao().Delete(data)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", record)
}

func (o *CategoriesCtr) SaveOrUpdate(r *ghttp.Request) {
	data := gfcommon.RequestPostData(r, g.Map{"title": "required", "slug": "required", "description": "required"})
	id := r.GetPostInt("id", 0)
	if id == 0 {
		_, e := dao.CategoriesDao().Save(data)
		response.JsonFailError(r, e, nil)
	} else {
		_, e := dao.CategoriesDao().Update(data, g.Map{"id": id})
		response.JsonFailError(r, e, nil)
	}
	response.JsonOK(r, "", nil)
}
