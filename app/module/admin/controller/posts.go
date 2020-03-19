package controller

import (
	"gf-app/app/dao"
	"gf-app/gfboot/gfcommon"
	"gf-app/library/auth"
	"gf-app/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

type PostsCtr struct{}

func init() {
	g.Server().BindObject("/admin/posts", new(PostsCtr))
}

func (o *PostsCtr) Get(r *ghttp.Request) {
	data := gfcommon.RequestGetData(r, g.Map{"id": "required"})
	record, e := dao.PostsDao().Find(data)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", record)
}

func (o *PostsCtr) List(r *ghttp.Request) {
	page := &gfcommon.Page{
		Limit: r.GetInt("limit", 10),
		Page:  r.GetInt("page", 1),
	}
	result, e := dao.PostsDao().Page(page, nil)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", result)
}

func (o *PostsCtr) Delete(r *ghttp.Request) {
	data := gfcommon.RequestGetData(r, g.Map{"id": "required"})
	record, e := dao.PostsDao().Delete(data)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", record)
}

func (o *PostsCtr) SaveOrUpdate(r *ghttp.Request) {
	data := gfcommon.RequestPostData(r, g.Map{"title": "required", "category_id": "required", "tags": "required", "content": "required", "htmlcontent": "required"})
	id := r.GetPostInt("id", 0)
	data["author_id"]=auth.GetUserId(r)
	if id == 0 {
		data["created_at"]=gtime.Now()
		_, e := dao.PostsDao().Save(data)
		response.JsonFailError(r, e, nil)
	} else {
		data["updated_at"]=gtime.Now()
		_, e := dao.PostsDao().Update(data, g.Map{"id": id})
		response.JsonFailError(r, e, nil)
	}
	response.JsonOK(r, "", nil)
}
