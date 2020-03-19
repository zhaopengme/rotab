package dao

import (
	"gf-app/gfboot/gfcommon"
	"gf-app/gfboot/gfdao"
	"github.com/gogf/gf/frame/g"
)

type postsDao struct {
	gfdao.Base
}

func PostsDao() *postsDao {
	d := new(postsDao)
	d.TableName = "rotab_posts"
	return d
}


func (this *postsDao) Page(page *gfcommon.Page, where interface{}) (*gfcommon.Page, error) {
	total, err := this.T().Where(where).Count()
	if err != nil {
		return page, err
	}
	w := g.Map{}
	if where != nil {
		for k, v := range where.(map[string]interface{}) {
			w["rp."+k] = v
		}
	} else {
		w = nil
	}
	data, err := this.DB().Table("rotab_posts rp").LeftJoin("rotab_users ru", "rp.author_id=ru.id").LeftJoin("rotab_categories rc","rp.category_id = rc.id").Fields("rp.*,ru.nickname,rc.title as ctitle").Where(w).Order("rp.id desc").Limit((page.Page-1)*page.Limit, page.Limit).All()
	if err != nil {
		page.Paginate(total, nil)
		return page, err
	}
	page.Paginate(total, data.List())
	return page, nil
}
