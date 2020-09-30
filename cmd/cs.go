package cmd

import "net"

func csHandler(c cmd, conn net.Conn, args []string) {
	if len(args) >= 2 {
		if args[1] == "help" {
			c.Help(conn)
			return
		}
	}
}
