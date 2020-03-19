package view

import (
	"gf-app/app/dao"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

func init() {
	glog.Debug("init view")
	v := g.View()
	v.BindFunc("renderTime", RenderTime)
	v.BindFunc("strSplit", gstr.Split)
	v.BindFunc("getNavs", GetNavs)
	v.BindFunc("themeFile", ThemeFile)
	v.BindFunc("getNextPost", GetNextPost)
	v.BindFunc("getPrevPost", GetPrevPost)
	v.BindFunc("getPageLink", GetPageLink)
	v.BindFunc("getComments", GetComments)
	v.BindFunc("mergeObj", MergeObj)
}

func RenderTime(start int64) string {
	return gconv.String(gtime.Now().Microsecond()) + "ms"
}

func ThemeFile(file string) string {
	return "themes/default/" + file
}

func CommonParams(r *ghttp.Request) {
	records, _ := dao.OptionDao().FindAll(g.Map{"category": "general"})
	params := g.Map{"r": r}
	for _, record := range records {
		params[record["category"].String()+gstr.UcFirst(record["name"].String())] = record["value"].String()
	}
	g.View().Assigns(params)
}

func GetNextPost(postId interface{}) interface{} {
	record, e := dao.PostsDao().T().Where("id>?", postId).OrderBy("id asc").Limit(1).One()
	if e != nil {
		return nil
	}
	return record.Map()
}

func GetPrevPost(postId interface{}) interface{} {
	record, e := dao.PostsDao().T().Where("id<?", postId).OrderBy("id desc").Limit(1).One()
	if e != nil {
		return nil
	}
	return record.Map()
}

func GetComments(postId interface{}, parentCommentId interface{}) interface{} {
	if len(gconv.String(postId)) == 0 {
		return nil
	}
	where := g.Map{"article_id": postId, "status": "ok"}
	where["parent_comment_id"] = parentCommentId
	if len(gconv.String(parentCommentId)) == 0 {
		where["parent_comment_id"] = 0
	}
	records, e := dao.CommentDao().T().Where(where).OrderBy("id desc").All()
	if e != nil {
		return nil
	}
	return records.List()
}

func MergeObj(args ...interface{}) interface{} {
	params := g.Map{}
	for k, v := range args {
		if k%2 == 0 {
			params[gconv.String(v)] = args[k+1]
		}
	}
	return params
}

func GetNavs() interface{} {
	records, e := dao.NavigationsDao().T().Where("status=?", "yes").OrderBy("number desc").All()
	if e != nil {
		return nil
	}
	return records.List()
}

func GetPageLink(page interface{}, params map[string]interface{}) string {
	r := params["r"].(*ghttp.Request)
	requestParams := params["Request"].(map[string]interface{})
	category := gconv.String(requestParams["category"])
	tag := gconv.String(requestParams["tag"])
	curPage := gconv.String(page)
	url := ""
	if len(category) > 0 {
		url = url + "/category/" + category
	}
	if len(tag) > 0 {
		url = url + "/tag/" + tag
	}
	if len(curPage) > 0 {
		url = url + "/page-" + curPage
	}
	if len(r.URL.RawQuery) > 0 {
		url = url + "?" + r.URL.RawQuery
	}
	return url
}
