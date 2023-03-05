package crypto

import (
	"fmt"
	"testing"
)

var newEncText = `JjPwapQGe22Xq+4ZnHr+A25zJZDgZD8U/YmQvRTfUvlDmhW9uIw+NgKupRVY5jQF8KA+CksAALccPd17e9t26fDGXgZdSn4SYijIcgnTcVcAPpxVdUt/T6Y2U9nycsmJHLH99GDKuvt9fuoJHqd9IhaiImd+xc3HB1V4EpAY6RQTMsw34VcFOf1jXG3VVOceGss1ZygXsX+uvcoml/bpzsI3zpmvk5YnNjSG36E1Y7KB5gY/zANA26XhsaHQF5ClrV2yrAyv+nR5wR7secDaKT/Xd6rXL7FOa7woplhnavVWrijjuROzn5ZG4oGOQt1eB3akSLHqbKiT6Fg24U3QTUtQh754i0V22JfnFqwsXEZs0DwlcK2dIWDEfL6+FVUlXZNUSzVdPQNl1g02pwbCgRwSzIiAO0swgxj3a0la2hQ0tGrn15P/oMf1yERw8apFX3k7eVnc8f8awDhFCULI+tFoprFObINPORKA324+azrsgLXsbjd2YJQZ9pttYq1/4thyKXR58hzYv5PGovepzuMs6q8CEIBJA7zzJJdB33Q=`

func TestLua(t *testing.T) {
	deData, err := Decrypt([]byte(newEncText))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(deData[:100])
	fmt.Println([]byte(`<root>
	 <story name="测试一下的剧情">
	<action type="DIALOG"`))
	enByte, err := Encrypt(deData)
	if err != nil {
		t.Fatal(err)
	}
	if string(enByte) != newEncText {
		t.Fatal("加密失败")
	}

}

func TestDecryptLuaFile(t *testing.T) {
	dir := `F:\Games\[PC]SSWS_HG_v21\gamedata\modcache\SSWS_HG\lua`
	dest := `F:\Games\[PC]SSWS_HG_v21\gamedata\modcache\SSWS_HG\lua_bak`
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
	dir := `F:\Games\[PC]SSWS_HG_v21\gamedata\modcache\SSWS_HG\Scripts`
	dest := `F:\Games\[PC]SSWS_HG_v21\gamedata\modcache\SSWS_HG\Scripts_bak`
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
