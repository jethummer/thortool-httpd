package main

import (
	"fmt"
	"log"
	"net/http"
)

var webPath string
var webPort string

func HttpServer() {

	fmt.Println(" ---------------------------------------------")
	fmt.Println(" |                   httpd                   |")
	fmt.Println(" ---------------------------------------------")
	fmt.Println(" |   Port: "+ webPort + "  |  Path: " + webPath)
	fmt.Println(" ---------------------------------------------")

	log.Fatal(http.ListenAndServe(":"+webPort, http.FileServer(http.Dir(webPath))))
}
