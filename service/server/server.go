package main

import (
	"flag"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/yaltachen/calendar/service"
)

var host string

func init() {
	flag.StringVar(&host, "host", "127.0.0.1:8000", "specify service host")
	flag.Parse()
	if host == "" {
		panic("must specify host")
	}
}

func main() {
	var (
		listener net.Listener
		conn     net.Conn
		err      error
	)
	err = rpc.Register(service.DateTransServer{})
	if err != nil {
		panic(err)
	}

	if listener, err = net.Listen("tcp", host); err != nil {
		panic(err)
	}

	for {
		if conn, err = listener.Accept(); err != nil {
			log.Printf("accept faild, err: %v", err)
		}
		log.Printf("got conn")
		go jsonrpc.ServeConn(conn)
	}
}
