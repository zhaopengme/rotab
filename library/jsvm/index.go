package jsvm

import (
	"github.com/dop251/goja"
)

func RunJsStr(jsstr string) (interface{}, error) {
	jsvm := goja.New()
	jsvm.Set("version", "rotabl")
	val, e := jsvm.RunString(jsstr)
	if e != nil {
		return nil, e
	}
	return val.Export(), nil
}
