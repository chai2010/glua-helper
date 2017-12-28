// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package helper

import (
	"fmt"
	"strconv"

	lua "github.com/yuin/gopher-lua"
)

func Check(L *lua.LState, idx0 int, args ...interface{}) {
	for i, a := range args {
		switch a := a.(type) {
		case *bool:
			*a = L.CheckBool(idx0 + i)
		case *int:
			*a = L.CheckInt(idx0 + i)
		case *string:
			*a = L.CheckString(idx0 + i)
		case []int:
			a = append(a, CheckIntList(L, idx0+i)...)
		case []string:
			a = append(a, CheckStringList(L, idx0+i)...)
		default:
			panic(fmt.Sprintf("unknown type: %T", a))
		}
	}
}

func CheckAnyList(L *lua.LState, n int) []lua.LValue {
	v := L.Get(n)
	if tb, ok := v.(*lua.LTable); ok {
		var ret []lua.LValue
		for i := 0; i < tb.Len(); i++ {
			ret = append(ret, tb.RawGetInt(i))
		}
		return ret
	}

	L.TypeError(n, lua.LTTable)
	return nil
}

func CheckIntList(L *lua.LState, n int) []int {
	v := L.Get(n)
	if tb, ok := v.(*lua.LTable); ok {
		var ret []int
		for i := 0; i < tb.Len(); i++ {
			item := tb.RawGetInt(i)
			if lv, ok := item.(lua.LNumber); ok {
				ret = append(ret, int(lv))
			} else {
				x, _ := strconv.Atoi(lv.String())
				ret = append(ret, x)
			}
		}
		return ret
	}

	L.TypeError(n, lua.LTTable)
	return nil
}

func CheckStringList(L *lua.LState, n int) []string {
	v := L.Get(n)
	if tb, ok := v.(*lua.LTable); ok {
		var ret []string
		for i := 0; i < tb.Len(); i++ {
			item := tb.RawGetInt(i)
			if lv, ok := item.(lua.LString); ok {
				ret = append(ret, string(lv))
			} else {
				ret = append(ret, item.String())
			}
		}
		return ret
	}

	L.TypeError(n, lua.LTTable)
	return nil
}
