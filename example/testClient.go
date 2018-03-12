package main

import (
	"flag"
	"gochat"
)

func main() {
	isClient := flag.Bool("p", false, "if is client platform true else false")
	flag.Parse()

	if *isClient {
		gochat.Client()
	} else {
		gochat.StartServer()
	}

}
