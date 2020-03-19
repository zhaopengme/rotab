package database

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	glog.Debug("init database")
	g.DB().SetDebug(g.Config().GetBool("database.default.debug"))
}
