package crypto

import (
	"fmt"
	"testing"
)

var newEncText = `VFOTUtI39pwaDS62Vscga3suoofLFJhJ+aBVCuRKHZAYh1zfQ9tysnFZOFaOHOV5NQCUq7Wq4RcY18sbjB7pD33jOUJ5j97r2T3PhPqOaelT4LlTj0O1DHOYSC3JqlMmIWs1z74BfhL3UAXLs1oiuEOXLeZ8QHakEN1ZSOtQqLKnGM2wjGsDb2dDYAMOyr9SnFNLvSa20KCz2tpcuvUvp9A4oNby1IzI4v0OnxBBqo7BbBiyZDqQRcfC2U0xyv3PuDLAMs6wYb8b/j8WJWeUD0H1ydOvCbnGn0rI8yIJ8ArCeEUoHcbwxyUQM4GrN7/Flix6o3aYxOrIRcjEswVkrwl1nE4qqgL5UmT+yIuhkcfxHsw6yIIApbCoc7dKsTPRTehoUZ44fBpCP7F46H+kYfqJ5JaP3/bE/CwSbzzZymMEDFaBWgcwvrc/OGCnF8ZUiyUi5EA++az212Yab7c8JwCUnU8uuPfTPMbTyiXaQnWgA3wmuyhUYQ6So8dyg9ewGc0XMsSn2e1XKAGUVDUQItZkblEY6WopLaAiZiq6zGDUyBkbWWix6r5hdVlYSiDmQsDVwrbE8qsLLe7/IKejsTNLSXi46W0EV3HggTRJLsY6nGr5ev2p0GwTXBX2UFEBObFEXdQDLQJSOEAvwY9pMkOK0Y9aro3TWNJ8+gvU5yNtUzCG0K6CAp4M5Xq6BVgZ`

func TestLua(t *testing.T) {
	deData, err := Decrypt([]byte(newEncText))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(deData))
	enByte, err := Encrypt(deData)
	if err != nil {
		t.Fatal(err)
	}
	if string(enByte) != newEncText {
		t.Fatal("加密失败")
	}
}

func TestDecryptLuaFile(t *testing.T) {
	dir := `F:\Games\JinX119\gamedata\modcache\SSWS_HG\lua_bak`
	dest := `F:\Games\JinX119\gamedata\modcache\SSWS_HG\lua_bak`
	err := RunDir(dir, dest, ModeDe, []string{".lua"})
	if err != nil {
		t.Fatalf("%+v", err)
	}
}

func TestEncryptLuaFile(t *testing.T) {
	dir := `F:\Games\JinX119all\gamedata\modcache\SSWS_HG\lua-out`
	dest := `F:\Games\JinX119all\gamedata\modcache\SSWS_HG\lua-out`
	err := RunDir(dir, dest, ModeEn, []string{".lua"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecryptXmlFile(t *testing.T) {
	dir := `H:\Games\金X无双后宫V111\gamedata\modcache\SSWS_HG\Scripts`
	dest := `H:\Games\金X无双后宫V111\mod\Scripts`
	err := RunDir(dir, dest, ModeDe, []string{".xml"})
	if err != nil {
		t.Fatal(err)
	}
}
func TestEncryptXmlFile(t *testing.T) {
	dir := `H:\Games\金X无双后宫V111\mod\Scripts`
	dest := `H:\Games\金X无双后宫V111\mod_en\Scripts`
	err := RunDir(dir, dest, ModeEn, []string{".xml"})
	if err != nil {
		t.Fatal(err)
	}
}
