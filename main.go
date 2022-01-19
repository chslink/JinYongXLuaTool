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
	mode = flag.String("mode", "d", "模式，e 表示加密模式,d 表示解密模式")
	src  = flag.String("src", "", "源文件夹，默认为当前文件夹")
	dest = flag.String("dest", "", "目标文件夹，默认为当前文件夹")
	ext  = flag.String("ext", ".lua,.xml", "扩展名，只有对应扩展名才会被加密解密")
)

func usage() {
	w := flag.CommandLine.Output()
	_, _ = fmt.Fprintf(w, "Usage of %s:  \n", os.Args[0])
	flag.PrintDefaults()
	_, _ = fmt.Fprintf(w,
		`例如： 加密lua文件 %s jxt.exe -mode e -src "H:\Game\JinYongX\gamedata\modcache\SSWS_HG\lua" -dest "H:\Game\JinYongX\gamedata\modcache\SSWS_HG\lua" %s`,
		"\n", "\n 或者直接将jxt.exe放进文件夹执行\n jxt.exe -mode d")
}

var (
	wd string
)

func main() {
	flag.Usage = usage
	flag.Parse()
	_wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	wd = _wd
	srcPath := getPath(*src)
	destPath := getPath(*dest)
	exts := strings.Split(*ext, ",")
	switch *mode {
	case "e":
		err = crypto.EncryptDir(srcPath, destPath, exts...)
	case "d":
		err = crypto.DecryptDir(srcPath, destPath, exts...)
	}
	if err != nil {
		log.Printf("程序执行错误\n")
		log.Print(err)
		fmt.Println()
		flag.Usage()
		os.Exit(1)
	} else {
		log.Println("操作完成")
	}
}

func getPath(path string) string {
	if path == "" {
		return wd
	}
	return filepath.Clean(path)
}
