package crypto

import (
	"testing"
)

func TestDecryptLuaFile(t *testing.T) {
	dir := `F:\Games\[PC]SSWS_HG_v21\gamedata\modcache\SSWS_HG\lua`
	dest := `F:\Games\[PC]SSWS_HG_v21\gamedata\modcache\SSWS_HG\lua_bak`
	err := RunDir(dir, dest, ModeDe, []string{".lua"}, ver122)
	if err != nil {
		t.Fatalf("%+v", err)
	}
}

func TestEncryptLuaFile(t *testing.T) {
	dir := `F:\Games\JinX119all\gamedata\modcache\SSWS_HG\lua-out`
	dest := `F:\Games\JinX119all\gamedata\modcache\SSWS_HG\lua-out`
	err := RunDir(dir, dest, ModeEn, []string{".lua"}, ver119)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDecryptXmlFile(t *testing.T) {
	dir := `D:\Games\JinX20230909\gamedata\modcache\SSWS\Scripts`
	dest := `D:\Games\JinX20230909\gamedata\modcache\SSWS\Scripts_bak`
	err := RunDir(dir, dest, ModeDe, []string{".xml"}, ver20230909)
	if err != nil {
		t.Fatal(err)
	}
}
func TestEncryptXmlFile(t *testing.T) {
	dir := `F:\Games\JinX119all\gamedata\modcache\SSWS_HG\Scripts-bak`
	dest := `F:\Games\JinX119all\gamedata\modcache\SSWS_HG\Scripts-bak`
	err := RunDir(dir, dest, ModeEn, []string{".xml"}, ver119)
	if err != nil {
		t.Fatal(err)
	}
}
