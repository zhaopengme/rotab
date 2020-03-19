package boot

import (
    _ "gf-app/boot/config"
    _ "gf-app/boot/database"
    _ "gf-app/boot/server"
    _ "gf-app/boot/view"
    "github.com/gogf/gf/os/glog"
)

func init() {
    glog.Debug("init boot")
}
