package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)

func Gowire(name string) string {
	return Invoke(name)
}

func funcHello(name string) string {
	return fmt.Sprintf(`Hello %s`, name)
}

func main() {
	g.View().BindFunc("gowire", Gowire)
	g.View().BindFunc("hello", funcHello)

	Register("hello", New)
	Register("demo", NewDemo)

	content := `aaaa {{gowire "hello"}} aaaa`
	result1, _ := g.View().ParseContent(content, nil)
	fmt.Println(result1)
}
