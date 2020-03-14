package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	help  bool
	path  string
	port  string
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")

	flag.StringVar(&port, "p", "8088", "web server port")
	flag.StringVar(&path, "w", "../Publish/Editor", "web root path")

	// 改变默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `httpd version: 0.0.1
Usage: httpd [-help] [-p port] [-w webpath]
`)
	flag.PrintDefaults()
}

func ParseArgs() {

	flag.Parse()

	if help || len(path) == 0 || len(port) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	webPort = port
	webPath = path

}
