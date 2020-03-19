package gfdao

import (
	"database/sql"
	"gf-app/gfboot/gfcommon"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

type Base struct {
	TableName string
}

func (this *Base) T() *gdb.Model {
	return g.DB().Table(this.TableName)
}
func (this *Base) DB() gdb.DB {
	return g.DB()
}
func (this *Base) Find(where interface{}) (gdb.Record, error) {
	return this.T().Where(where).One()
}

func (this *Base) FindAll(where interface{}) (gdb.Result, error) {
	return this.T().Where(where).All()
}

func (this *Base) FindFieldsAll(fields string, where interface{}) (gdb.Result, error) {
	return this.T().Fields(fields).Where(where).All()
}

func (this *Base) Save(data interface{}) (sql.Result, error) {
	return this.T().Data(data).Insert()
}

func (this *Base) Delete(where interface{}) (sql.Result, error) {
	return this.T().Where(where).Delete()
}

func (this *Base) Update(data interface{}, where interface{}) (sql.Result, error) {
	return this.T().Data(data).Where(where).Update()
}

func (this *Base) Page(page *gfcommon.Page, where interface{}) (*gfcommon.Page, error) {
	total, err := this.T().Where(where).Count()
	if err != nil {
		return page, err
	}
	data, err := this.T().Where(where).Limit((page.Page-1)*page.Limit, page.Limit).All()
	if err != nil {
		page.Paginate(total, nil)
		return page, err
	}
	page.Paginate(total, data.List())
	return page, nil
}
