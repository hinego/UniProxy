package main

import (
	"UniProxy/proxy"
	"UniProxy/router"
	"flag"
	"log"
	"strconv"
)

var host = flag.String("host", "127.0.0.1", "host")
var port = flag.Int("port", 33212, "port")

func main() {
	flag.Parse()
	proxy.ResUrl = "http://127.0.0.1:" + strconv.Itoa(*port)
	router.Init()
	log.Println("UniProxy start at", *host, *port)
	if err := router.Start(*host, *port); err != nil {
		log.Fatalln("start error:", err)
	}
}
