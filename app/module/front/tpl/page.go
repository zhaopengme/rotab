package tpl

import (
	"gf-app/app/dao"
	"gf-app/boot/view"
	"gf-app/gfboot/gfcommon"
	"gf-app/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/os/gview"
	"net/http"
)

type Page struct {
}

func init() {
	p := &Page{}
	s := g.Server()

	s.BindHandler("/", p.Index)
	s.BindHandler("/page-{page}", p.Index)
	s.BindHandler("/comment", p.Comment)
	s.BindHandler("/themes/*/**", p.Themes)
	s.BindHandler("/article/{id}", p.Article)
	s.BindHandler("/category/{category}", p.Index)
	s.BindHandler("/category/{category}/page-{page}", p.Index)
	s.BindHookHandler("/", ghttp.HOOK_BEFORE_SERVE, view.CommonParams)
	s.BindHookHandler("/category/*/**", ghttp.HOOK_BEFORE_SERVE, view.CommonParams)
	s.BindHookHandler("/article/*/**", ghttp.HOOK_BEFORE_SERVE, view.CommonParams)
}

func (p *Page) Index(r *ghttp.Request) {
	page := &gfcommon.Page{
		Limit: r.GetInt("limit", 10),
		Page:  r.GetInt("page", 1),
	}
	where := g.Map{}
	category := r.Get("category", "")
	if category != "" {
		where["category_id"] = category
	}
	result, e := dao.PostsDao().Page(page, where)
	response.JsonFailError(r, e, nil)
	r.Response.WriteTpl("themes/default/index.html", gview.Params{"posts": result})
}

func (p *Page) Article(r *ghttp.Request) {
	id := r.GetInt("id", 0)
	if id == 0 {
		r.Response.Status = http.StatusNotFound
		r.ExitAll()
	}
	result, e := dao.PostsDao().T().Where(g.Map{"id": id}).One()
	response.JsonFailError(r, e, nil)

	r.Response.WriteTpl("themes/default/post.html", gview.Params{"post": result.Map()})
}

func (p *Page) Themes(r *ghttp.Request) {
	staticFile := gfile.Join(gfile.RealPath(""), "template", r.URL.Path)
	if gfile.Exists(staticFile) {
		r.Response.ServeFile(staticFile)
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

func (p *Page) Comment(r *ghttp.Request) {
	data := gfcommon.RequestPostData(r, g.Map{"article_id": "required", "author_email": "required", "content": "required", "author_name": "required"})
	data["user_agent"] = r.UserAgent()
	data["ip"] = UserIP(r.Request)
	data["created_at"] = gtime.Now()
	data["status"] = "waiting"
	data["author_id"] = r.GetInt("author_id", 0)
	data["parent_comment_id"] = r.GetInt("parent_comment_id", 0)
	data["author_url"] = r.GetString("author_url", "")

	dao.CommentDao().Save(data)

	r.Cookie.Set("author_email", data["author_email"].(string))
	r.Cookie.Set("author_name", data["author_name"].(string))
	r.Cookie.Set("author_url", data["author_url"].(string))

	pageUrl := r.GetPostString("page_url", "/")
	r.Response.RedirectTo(pageUrl)
}

func UserIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}
