package crypto

type VersionConfig struct {
	Version string // 版本名词
	// 默认解密密钥
	Key string
	IV  string
	// 解密Lua脚本密钥 默认为空 如果配置该字段 解密Lua会优先使用
	LuaKey     string
	LuaIV      string
	FixLua     bool
	FixLuaFunc func(filename string, data []byte) ([]byte, error)
	// 解密XML脚本密钥 默认为空 如果配置字段，优先使用
	XmlKey     string
	XmlIV      string
	FixXml     bool
	FixXmlFunc func(filename string, data []byte) ([]byte, error)
}

var ConfigList = []*VersionConfig{ver111, ver119, ver122, ver20230909}
var ConfigMap = func() map[string]*VersionConfig {
	var configMap = make(map[string]*VersionConfig)
	for _, v := range ConfigList {
		configMap[v.Version] = v
	}
	return configMap
}()
var DefaultConfig = ver20230909
