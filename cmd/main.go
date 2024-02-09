package main

import "url-shortener/internal/net"

func main() {
	app := net.Init()
	app.Run()
}
