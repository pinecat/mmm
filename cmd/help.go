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
		conn.Write([]byte("Help menu coming soon.\n"))
	},
}
