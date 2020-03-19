package main

import (
	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(`print("hello")`); err != nil {
		panic(err)
	}
	if err := L.DoFile(`/Users/zhaopeng/Projects/user/lua/module.lua`); err != nil {
		panic(err)
	}

}


