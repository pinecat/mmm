package serve

import (
	"fmt"
	"net"
	"net/textproto"
)

func Start() {
	ln, err := net.Listen("tcp", ":25578")
	if err != nil {
		fmt.Printf("%v", err)
	}
	for {
		conn, _ := ln.Accept()
		tpconn := textproto.NewConn(conn)
		go func() {
			for {
				data, _ := tpconn.ReadLine()
				fmt.Printf("[%s] %s.\n", conn.RemoteAddr().String(), string(data))
			}
		}()
	}
}
