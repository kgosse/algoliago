package main

import "flag"

func main() {
	var port string
	flag.StringVar(&port, "p", "8080", "server port")
	flag.Parse()
	serve(port)
}
