package main

import (
	"flag"
)

func main() {
	conf := flag.String("c", ".", "configuration file")

	flag.Parse()

	serv := newPiServer(conf)
	serv.Start()
}
