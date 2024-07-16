package crypto

import (
	"strings"
	"testing"
)

var var122EncText = `JjPwapQGe22Xq+4ZnHr+A25zJZDgZD8U/YmQvRTfUvlDmhW9uIw+NgKupRVY5jQF8KA+CksAALccPd17e9t26fDGXgZdSn4SYijIcgnTcVcAPpxVdUt/T6Y2U9nycsmJHLH99GDKuvt9fuoJHqd9IhaiImd+xc3HB1V4EpAY6RQTMsw34VcFOf1jXG3VVOceGss1ZygXsX+uvcoml/bpzsI3zpmvk5YnNjSG36E1Y7KB5gY/zANA26XhsaHQF5ClrV2yrAyv+nR5wR7secDaKT/Xd6rXL7FOa7woplhnavVWrijjuROzn5ZG4oGOQt1eB3akSLHqbKiT6Fg24U3QTUtQh754i0V22JfnFqwsXEZs0DwlcK2dIWDEfL6+FVUlXZNUSzVdPQNl1g02pwbCgRwSzIiAO0swgxj3a0la2hQ0tGrn15P/oMf1yERw8apFX3k7eVnc8f8awDhFCULI+tFoprFObINPORKA324+azrsgLXsbjd2YJQZ9pttYq1/4thyKXR58hzYv5PGovepzuMs6q8CEIBJA7zzJJdB33Q=`

func TestXML122(t *testing.T) {
	deData, err := Decrypt(ver122, "test.xml", []byte(var122EncText))
	if err != nil {
		t.Fatal(err)
	}
	str := strings.ReplaceAll(string(deData), "\r\n", "\n")
	text := `<root>
  <story name="测试一下的剧情">
<action type="DIALOG" value="主角#这是BUG吗？卧槽这绝对是bug吧？？" />
<action type="DIALOG" value="主角#这是BUG吗？卧槽这绝对是bug吧？？" />
<action type="DIALOG" value="主角#这是BUG吗？卧槽这绝对是bug吧？？" />
<action type="DIALOG" value="主角#这是BUG吗？卧槽这绝对是bug吧？？" />
</story>
</root>`
	if str != text {
		t.Fatal("解密失败")
	}
	//enByte, err := Encrypt(deData)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if string(enByte) != var122EncText {
	//	t.Fatal("加密失败")
	//}

}
