package cmd

import (
	"net"
)

var cmdHelp cmd = cmd{
	Name:        "help",
	Aliases:     []string{"h"},
	Description: "Print a general help menu.",
	Type:        "Command",
	Usage:       "help",
	Example:     "help",
	SubCmds:     []cmd{},
	Handler: func(c cmd, conn net.Conn, args []string) {
		if len(args) >= 2 {
			if args[1] == "help" {
				c.Help(conn)
				return
			}
		}
		conn.Write([]byte("Help menu coming soon.\n"))
	},
}
