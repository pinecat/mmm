package cmd

import "net"

var cmdJava cmd = cmd{
	Name:        "java",
	Aliases:     []string{"j"},
	Description: "Checks if java is installed on the system.",
	Type:        "Command",
	Usage:       "java",
	Example:     "java",
	SubCmds:     []cmd{},
	Handler: func(c cmd, conn net.Conn, args []string) {
		conn.Write([]byte("[mmm] Not implemented yet ¯\\_(ツ)_/¯.\n"))
	},
}
