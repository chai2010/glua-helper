// Copyright 2017 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package helper

import (
	"strings"

	lua "github.com/yuin/gopher-lua"
)

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
