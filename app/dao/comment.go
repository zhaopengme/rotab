package dao

import (
	"gf-app/gfboot/gfcommon"
	"gf-app/gfboot/gfdao"
	"github.com/gogf/gf/frame/g"
)

type commentDao struct {
	gfdao.Base
}

func CommentDao() *commentDao {
	d := new(commentDao)
	d.TableName = "rotab_comments"
	return d
}

func (this *commentDao) Page(page *gfcommon.Page, where interface{}) (*gfcommon.Page, error) {
	total, err := this.T().Where(where).Count()
	if err != nil {
		return page, err
	}
	w := g.Map{}
	if where != nil {
		for k, v := range where.(map[string]interface{}) {
			w["rc."+k] = v
		}
	} else {
		w = nil
	}
	data, err := this.DB().Table("rotab_comments rc").LeftJoin("rotab_posts rp", "rc.article_id=rp.id").Fields("rc.*,rp.title").Where(w).Limit((page.Page-1)*page.Limit, page.Limit).All()
	if err != nil {
		page.Paginate(total, nil)
		return page, err
	}
	page.Paginate(total, data.List())
	return page, nil
}
