package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/chslink/JinYongXLuaTool/crypto"
)

var (
	mode = flag.String("mode", "", "模式，e 表示加密模式,d 表示解密模式")
	src  = flag.String("src", "", "源文件夹，默认为当前文件夹")
	dest = flag.String("dest", "", "目标文件夹，默认为当前文件夹")
	ver  = flag.String("ver", "", "游戏版本，当前只支持111 119,默认119")
	key  = flag.String("key", "", "自定义解密key")
	iv   = flag.String("iv", "", "自定义解密iv")
	ext  = flag.String("ext", ".lua,.xml", "扩展名，只有对应扩展名才会被加密解密")
)

func usage() {
	w := flag.CommandLine.Output()
	_, _ = fmt.Fprintf(w, "Usage of %s:  \n", os.Args[0])
	flag.PrintDefaults()
	_, _ = fmt.Fprintf(w,
		`例如： 加密lua文件 %s jxt.exe -mode e -src "H:\Game\JinYongX\gamedata\modcache\SSWS_HG\lua" -dest "H:\Game\JinYongX\gamedata\modcache\SSWS_HG\lua" %s`,
		"\n", "\n 或者直接将jxt.exe放进文件夹执行")
}

var (
	wd string
)

func main() {
	flag.Usage = usage
	flag.Parse()
	curVer := crypto.DefaultConfig
	if *key != "" && *iv != "" {
		curVer.Key = *key
		curVer.IV = *iv
	} else {
		if *ver != "" {
			v, ok := crypto.ConfigMap[*ver]
			if ok {
				curVer = v
			}
		}
	}

	_wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	wd = _wd
	srcPath := getPath(*src)
	destPath := getPath(*dest)
	splitExt := strings.Split(*ext, ",")
	err = crypto.RunDir(srcPath, destPath, *mode, splitExt, curVer)
	if err != nil {
		log.Printf("程序执行错误:\n%+v\n", err)
		flag.Usage()
	} else {
		log.Println("操作完成")
	}
	fmt.Println("按任意键继续")
	fmt.Scanf("按任意键继续")
}

func getPath(path string) string {
	if path == "" {
		return wd
	}
	return filepath.Clean(path)
}
