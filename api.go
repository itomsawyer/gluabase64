package gluabase64

import (
	"bytes"
	"encoding/base64"
	"io"

	glua "github.com/yuin/gopher-lua"
)

var api = map[string]glua.LGFunction{
	"decode": decode,
	"encode": encode,
}

func decode(L *glua.LState) int {
	var enc *base64.Encoding

	str := L.CheckString(1)
	if L.GetTop() > 1 {
		encStr := L.CheckString(2)
		if enc = encoding(encStr); enc == nil {
			L.Push(glua.LNil)
			L.Push(glua.LString("encoding not support"))
			return 2
		}
	}

	if enc == nil {
		enc = base64.URLEncoding
	}

	r := bytes.NewBufferString(str)
	w := new(bytes.Buffer)
	d := base64.NewDecoder(enc, r)
	if _, err := io.Copy(w, d); err != nil {
		L.Push(glua.LNil)
		L.Push(glua.LString(err.Error()))
		return 2
	}

	L.Push(glua.LString(w.String()))
	L.Push(glua.LNil)
	return 2
}

func encode(L *glua.LState) int {
	var enc *base64.Encoding

	str := L.CheckString(1)
	if L.GetTop() > 1 {
		encStr := L.CheckString(2)
		if enc = encoding(encStr); enc == nil {
			L.Push(glua.LNil)
			L.Push(glua.LString("encoding not support"))
			return 2
		}
	}

	if enc == nil {
		enc = base64.URLEncoding
	}

	r := bytes.NewBuffer([]byte(str))

	w := new(bytes.Buffer)
	e := base64.NewEncoder(enc, w)
	if _, err := io.Copy(e, r); err != nil {
		L.Push(glua.LNil)
		L.Push(glua.LString(err.Error()))
		return 2
	}

	e.Close()
	L.Push(glua.LString(w.String()))
	L.Push(glua.LNil)
	return 2
}

func encoding(enc string) *base64.Encoding {
	switch enc {
	case "raw_url":
		return base64.RawURLEncoding
	case "raw_std":
		return base64.RawStdEncoding
	case "std":
		return base64.StdEncoding
	case "url":
		return base64.URLEncoding
	}

	return nil
}
