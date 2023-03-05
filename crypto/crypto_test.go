package crypto

import (
	"testing"
)

var newEncText = `X5x9nvhLSHp5Sf5pOxYRDDPV4tfYFk5gbDGkQV7NywhXXL2jylwHq6xylGnTIN4FLjeu3jA+YEEP5XSD4em6LMCwODmOVK7MGz2kyi5PFDlUMHxcbgha8DsTVvB6SQHbEWix8FH0B3CSINIpNLHFPpXzRc09h58hMlaBl2v+JXm7DAxtYv92SEYG23XC7qfx0JnExlJPmCgBkvEnNT4K0Ce+1DPYhgLD1glHXfgYcKkN42OxUnadBjGJ6c7ideNyHJizCNMTXSjtT5dReVzV40qIXtNsbAU7GIDzIeCIrAdpoixigcuBcyiVKKUF+luDRN5beIgiyVGB2mM2NdOsmxq3iMPPrp7y02TVz4ntt9Y9SrTQDZwZOh/V1cEkSHfY1gK1iDbOEi80KRUrCe7SjTrbIFL0lj619jJYOWiBx+haAMjouJGnRjn5M0wzFtmV4jQViLWxX5mILPfeBpYQ6Z/cRecJALXAW1asoN/Hw4ygrt/mhqfuxWify91tVQippEp6gDXY1Dvk2E6q9dXxb8TG5IbxV8q6Xj/llBY2k3M=`

func TestLua(t *testing.T) {
	deData, err := Decrypt([]byte(newEncText))
	if err != nil {
		t.Fatal(err)
	}
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
	dir := `F:\Games\JinX119all\gamedata\modcache\SSWS_HG\Scripts`
	dest := `F:\Games\JinX119all\gamedata\modcache\SSWS_HG\Scripts-bak`
	err := RunDir(dir, dest, ModeDe, []string{".xml"})
	if err != nil {
		t.Fatal(err)
	}
}
func TestEncryptXmlFile(t *testing.T) {
	dir := `F:\Games\JinX119all\gamedata\modcache\SSWS_HG\Scripts-bak`
	dest := `F:\Games\JinX119all\gamedata\modcache\SSWS_HG\Scripts-bak`
	err := RunDir(dir, dest, ModeEn, []string{".xml"})
	if err != nil {
		t.Fatal(err)
	}
}
