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
	split string
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")

	flag.StringVar(&path, "p", "", "web server port, default is 8088")
	flag.StringVar(&port, "w", "", "web server filepath, default is ../Assets/StreamingAssets")

	// 改变默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `httpd version: 0.0.1
Usage: httpd [-help] [-p port] [-w watchpath]
`)
	flag.PrintDefaults()
}

func ParseArgs() {

	flag.Parse()

	if help {
		flag.Usage()
		return
	} else if path != "" {
		webPath = path
	} else {
		webPath = "../Assets/StreamingAssets"
	}
	if port != "" {
		webPort = port
	} else {
		webPort = "8088"
	}

}
