package main

import (
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"reflect"
)

var (
	instances = gmap.NewStrAnyMap(true)
)

func Register(name string, component interface{}) {
	instances.Set(name, component)
}

func Invoke(name string) string {
	f := instances.Get(name)
	value := reflect.ValueOf(f)
	if value.Kind() != reflect.Func {
		fmt.Println("function is not function")
		return ""
	}
	argValues := make([]reflect.Value, 0)
	values := value.Call(argValues)
	if len(values) > 0 {
		c := values[0]
		m := c.MethodByName("Render")
		renderArgs := make([]reflect.Value, 0)
		renderArgs = append(renderArgs, reflect.ValueOf("heaaa"))
		vv := m.Call(renderArgs)
		return vv[0].Interface().(string)
	}
	return ""
}
