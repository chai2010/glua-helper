// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package helper

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func Return(L *lua.LState, args ...interface{}) int {
	for _, a := range args {
		switch a := a.(type) {
		case error:
			L.Push(lua.LString(a.Error()))
		case bool:
			RetBool(L, a)
		case int:
			RetInt(L, a)
		case []int:
			RetIntList(L, a)
		case string:
			RetString(L, a)
		case []string:
			RetStringList(L, a)
		default:
			RetString(L, fmt.Sprint(a))
		}
	}
	return len(args)
}

func RetError(L *lua.LState, err error) int {
	if err != nil {
		L.Push(lua.LString(err.Error()))
	} else {
		L.Push(lua.LNil)
	}
	return 1
}

func RetBool(L *lua.LState, v bool) int {
	L.Push(lua.LBool(v))
	return 1
}

func RetInt(L *lua.LState, v int) int {
	L.Push(lua.LNumber(v))
	return 1
}

func RetIntList(L *lua.LState, vs []int) int {
	tb := L.NewTable()
	for _, v := range vs {
		tb.Append(lua.LNumber(v))
	}
	L.Push(tb)
	return 1
}

func RetString(L *lua.LState, v string) int {
	L.Push(lua.LString(v))
	return 1
}

func RetStringList(L *lua.LState, vs []string) int {
	tb := L.NewTable()
	for _, v := range vs {
		tb.Append(lua.LString(v))
	}
	L.Push(tb)
	return 1
}
