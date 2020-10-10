package cmd

import "net"

var cmdCreate cmd = cmd{
	Name:        "create",
	Aliases:     []string{"c"},
	Description: "Send create request to mmm daemon.",
	Type:        "Command",
	Usage:       "create ...",
	Example:     "create server 1.16.2",
	SubCmds:     []cmd{cmdCreateServer},
	Handler: func(c cmd, conn net.Conn, args []string) {
		if len(args) >= 2 {
			if args[1] == "help" {
				c.Help(conn)
				return
			}
		}
	},
}

var cmdCreateServer cmd = cmd{
	Name:        "server",
	Aliases:     []string{"s"},
	Description: "Create and deploy a Minecraft server instance.",
	Type:        "SubCommand",
	Usage:       "create server [version]",
	Example:     "cs 1.16.2",
	SubCmds:     []cmd{},
	Handler: func(c cmd, conn net.Conn, args []string) {
		if len(args) >= 2 {
			if args[1] == "help" {
				c.Help(conn)
				return
			}
		}
	},
}
