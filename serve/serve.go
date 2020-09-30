package serve

import (
	"fmt"
	"net"
	"net/textproto"
)

func Start(port string) {
	ln, err := net.Listen("tcp", ":"+port)
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
