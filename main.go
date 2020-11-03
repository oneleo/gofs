package main

import (
	"flag"
	"fmt"
	"os"
)

// https://blog.wu-boy.com/2017/02/write-command-line-in-golang/
// https://medium.com/@zamhuang/golang-%E5%A6%82%E4%BD%95%E8%AE%80%E5%8F%96-command-line-argument-flag-%E5%BF%85%E7%9F%A5%E7%9A%84%E5%B9%BE%E7%A8%AE%E7%94%A8%E6%B3%95-7dee79763a9e
var (
	port int64
	path string
)

func init() {
	flag.Int64Var(&port, "port", 80, "設置使用的 port")
	flag.Int64Var(&port, "p", 80, "設置使用的 port")
	flag.StringVar(&path, "path", "./", "設置根路徑")
	flag.StringVar(&path, "r", "./", "設置根路徑")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "使用方法：gofs [options] [root]\n")
		fmt.Fprintf(os.Stderr, "有 2 個參數需要設置：\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "範例：gofs -p 8081 -r \"./www\"")
	}
}

func main() {
	flag.Parse()
}
