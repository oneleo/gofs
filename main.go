package main

import (
	"flag"
	"fmt"
	"os"
)

// https://blog.wu-boy.com/2017/02/write-command-line-in-golang/
// https://medium.com/@zamhuang/golang-%E5%A6%82%E4%BD%95%E8%AE%80%E5%8F%96-command-line-argument-flag-%E5%BF%85%E7%9F%A5%E7%9A%84%E5%B9%BE%E7%A8%AE%E7%94%A8%E6%B3%95-7dee79763a9e
// https://darjun.github.io/2020/01/10/godailylib/flag/
var (
	port   int64
	folder string
)

func init() {
	const (
		defaultPort   = 80
		defaultFolder = "./"
		usagePort     = "Set listening port"
		usageFolder   = "Set shared folder"
	)
	//flag.Int64Var(&port, "port", defaultPort, usagePort)
	flag.Int64Var(&port, "p", defaultPort, usagePort+" (shorthand)")
	//flag.StringVar(&folder, "folder", defaultFolder, usageFolder)
	flag.StringVar(&folder, "f", defaultFolder, usageFolder+" (shorthand)")

	flag.Usage = func() {
		// example: go help doc
		fmt.Fprintf(os.Stderr, "Gofs is a small file server for Browsing local folder and files.\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\n\tgofs [-p port] [-f folder]\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  gofs\n\tStart this file server and listen on port 80 for browse \"./\" folder.\n")
		fmt.Fprintf(os.Stderr, "  gofs -p 8081 -r ./www\n\tStart this file server and listen on port 8081 for browse \"./www\" folder.")
	}
}

func main() {
	flag.Parse()
}
