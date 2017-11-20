package gluabase64

import (
	"testing"

	glua "github.com/yuin/gopher-lua"
)

func TestInit(t *testing.T) {
	const str = `
	local b64 = require("base64")
	assert(type(b64) == "table")
	assert(type(b64.decode) == "function")
	assert(type(b64.encode) == "function")
	`
	s := glua.NewState()
	s.PreloadModule("base64", Loader)
	if err := s.DoString(str); err != nil {
		t.Error(err)
	}

}

func TestEncoding(t *testing.T) {
	const str = `
	local b64 = require("base64")
	local str = "tomsawyer.me!"
	local data, err = b64.encode(str)
	assert(err == nil)
	assert(data == "dG9tc2F3eWVyLm1lIQ==")
	`
	s := glua.NewState()
	s.PreloadModule("base64", Loader)
	if err := s.DoString(str); err != nil {
		t.Error(err)
	}

}

func TestDecoding(t *testing.T) {
	const str = `
	local b64 = require("base64")
	local str = "dG9tc2F3eWVyLm1lIQ=="
	local data, err = b64.decode(str)
	assert(err == nil)
	assert(data == "tomsawyer.me!")
	`
	s := glua.NewState()
	s.PreloadModule("base64", Loader)
	if err := s.DoString(str); err != nil {
		t.Error(err)
	}
}
