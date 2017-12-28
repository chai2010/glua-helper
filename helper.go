// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package helper

import (
	"fmt"
	"log"
	"strings"

	lua "github.com/yuin/gopher-lua"
)

func MakeIntList(L *lua.LState, args ...int) *lua.LTable {
	tb := L.NewTable()
	for _, v := range args {
		tb.Append(lua.LNumber(v))
	}
	return tb
}

func MakeStringList(L *lua.LState, args ...string) *lua.LTable {
	tb := L.NewTable()
	for _, v := range args {
		tb.Append(lua.LString(v))
	}
	return tb
}

func MakeArray(L *lua.LState, args ...interface{}) *lua.LTable {
	tb := L.NewTable()
	for i := 0; i < len(args); i++ {
		if args[i] == nil {
			tb.Append(lua.LNil)
			continue
		}
		switch a := args[i].(type) {
		case bool:
			tb.Append(lua.LBool(a))
		case int:
			tb.Append(lua.LNumber(a))
		case []int:
			tb.Append(MakeIntList(L, a...))
		case string:
			tb.Append(lua.LString(a))
		case []string:
			tb.Append(MakeStringList(L, a...))
		default:
			tb.Append(lua.LString(fmt.Sprintf("<unsupport type: %T>", a)))
		}
	}
	return tb
}

func DoScript(L *lua.LState, path, content string) error {
	// dofile
	if path != "" && content == "" {
		return L.DoFile(path)
	}

	// dostring with name
	if path == "" {
		path = "<string>"
	}
	fn, err := L.Load(strings.NewReader(content), path)
	if err != nil {
		return err
	}

	L.Push(fn)
	return L.PCall(0, lua.MultRet, nil)
}

func Run(L *lua.LState, code string, args ...interface{}) []interface{} {
	L.SetGlobal("arg", MakeArray(L, args...))
	if err := DoScript(L, "", code); err != nil {
		log.Fatal(err)
	}

	values := make([]interface{}, L.GetTop())
	for idx := 1; idx <= len(values); idx++ {
		lv := L.Get(idx)
		if lv == lua.LNil {
			values[idx-1] = nil
			continue
		}

		switch lv.(type) {
		case lua.LBool:
			values[idx-1] = bool(lv.(lua.LBool))
		case lua.LNumber:
			values[idx-1] = int(lv.(lua.LNumber))
		case lua.LString:
			values[idx-1] = string(lv.(lua.LString))
		default:
			values[idx-1] = lv.String()
		}
	}

	return values
}
