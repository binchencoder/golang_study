package main

import (
	"fmt"
	"log"
	"time"
)

type serverOption struct {
	address string
	port    int
	path    string
	timeout time.Duration
	log     *log.Logger
}

func newOption() *serverOption {
	return &serverOption{
		address: "0.0.0.0",
		port:    8080,
		path:    "/var/test",
		timeout: time.Second * 5,
		log:     nil,
	}
}

func server(option *serverOption) {
	fmt.Printf("option: %#v\n", option)
}

func main() {
	opt := newOption()
	opt.port = 8585

	server(opt)
}
