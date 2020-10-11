package cmd

import (
	"net"
)

var cmdRemove cmd = cmd{
	Name:        "remove",
	Aliases:     []string{"r"},
	Description: "Send remove request to mmm daemon.",
	Type:        "Command",
	Usage:       "remove ...",
	Example:     "create server srv0",
	SubCmds:     []cmd{cmdRemoveServer},
	Handler: func(c cmd, conn net.Conn, args []string) {
		// This command on its own does not actually do anything
		conn.Write([]byte("[mmm] The remove command must be used with a subcommand.\n"))
	},
}

var cmdRemoveServer cmd = cmd{
	Name:        "server",
	Aliases:     []string{"s"},
	Description: "Remove a Minecraft server instance.",
	Type:        "SubCommand",
	Usage:       "remove server [version]",
	Example:     "rs srv0",
	SubCmds:     []cmd{},
	Handler: func(c cmd, conn net.Conn, args []string) {
		conn.Write([]byte("[mmm] Not implemented yet ¯\\_(ツ)_/¯.\n"))
	},
}
