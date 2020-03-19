package controller

import (
	"gf-app/app/dao"
	"gf-app/gfboot/gfcommon"
	"gf-app/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

type CommentsCtr struct{}

func init() {
	g.Server().BindObject("/admin/comments", new(CommentsCtr))
}

func (o *CommentsCtr) Get(r *ghttp.Request) {
	data := gfcommon.RequestGetData(r, g.Map{"id": "required"})
	record, e := dao.CommentDao().Find(data)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", record)
}

func (o *CommentsCtr) List(r *ghttp.Request) {
	where := gfcommon.RequestGetData(r, g.Map{"status": "required"})
	page := &gfcommon.Page{
		Limit: r.GetInt("limit", 10),
		Page:  r.GetInt("page", 1),
	}
	result, e := dao.CommentDao().Page(page, where)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", result)
}

func (o *CommentsCtr) Delete(r *ghttp.Request) {
	data := gfcommon.RequestGetData(r, g.Map{"id": "required"})
	record, e := dao.CommentDao().Delete(data)
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", record)
}

func (o *CommentsCtr) SaveOrUpdate(r *ghttp.Request) {
	data := gfcommon.RequestPostData(r, g.Map{"status": "required"})
	id := r.GetPostInt("id", 0)
	if id == 0 {
		data["created_at"]=gtime.Now()
		_, e := dao.CommentDao().Save(data)
		response.JsonFailError(r, e, nil)
	} else {
		data["updated_at"]=gtime.Now()
		_, e := dao.CommentDao().Update(data, g.Map{"id": id})
		response.JsonFailError(r, e, nil)
	}
	response.JsonOK(r, "", nil)
}
