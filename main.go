package main

import (
	"fmt"
	_ "gf-app/boot"
	_ "gf-app/router"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
)

func main() {


	fmt.Println(gmd5.EncryptString("111111" + "12345"))
	config := g.Config().GetMap("server")
	s := g.Server()
	s.SetConfigWithMap(config)
	s.Run()



}
