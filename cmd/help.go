package cmd

import (
	"net"
)

func helpHandler(c cmd, conn net.Conn, args []string) {
	if len(args) >= 2 {
		if args[1] == "help" {
			c.Help(conn)
			return
		}
	}
	conn.Write([]byte("Help menu coming soon.\n"))
}
