package jsvm

import (
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"testing"
)

func TestInit(t *testing.T) {
	vm := goja.New()
	new(require.Registry).Enable(vm)
	console.Enable(vm)

	loop := &EventLoop{
		vm:      vm,
		jobChan: make(chan func()),
	}

	new(require.Registry).Enable(vm)
	console.Enable(vm)
	vm.Set("setTimeout", loop.setTimeout)
	vm.Set("setInterval", loop.setInterval)
	vm.Set("clearTimeout", loop.clearTimeout)
	vm.Set("clearInterval", loop.clearInterval)

	polyfill_source := ghttp.GetContent("https://cdn.jsdelivr.net/npm/promise-polyfill@8/dist/polyfill.min.js")
	vue_source := ghttp.GetContent("https://cdn.jsdelivr.net/npm/vue/dist/vue.js")
	renderer_source := ghttp.GetContent("https://raw.githubusercontent.com/vuejs/vue/dev/packages/vue-server-renderer/basic.js")
	app_source := gfile.GetContents("/Users/zhaopeng/Documents/Rotab/library/jsvm/app.js")

	v, err := vm.RunString(`var process = { env: { VUE_ENV: "server", NODE_ENV: "production" }}; this.global = { process: process };`)
	if err != nil {
		panic(err)
	}
	v, err = vm.RunString(polyfill_source)
	if err != nil {
		panic(err)
	}
	v, err = vm.RunString(vue_source)
	if err != nil {
		panic(err)
	}
	v, err = vm.RunString(renderer_source)
	if err != nil {
		panic(err)
	}
	v, err = vm.RunString(app_source)
	if err != nil {
		panic(err)
	}
	println(v)
}





func TestRunJsStr(t *testing.T) {
	js:=`version`
	v,e:=RunJsStr(js)
	println(e)
	println(v.(string))
}
