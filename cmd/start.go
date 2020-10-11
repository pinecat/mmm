package cmd

import (
	"fmt"
	"net"

	"github.com/pinecat/mmm/instance"
)

var cmdStart cmd = cmd{
	Name:        "start",
	Aliases:     []string{"s"},
	Description: "Start a server instance.",
	Type:        "Command",
	Usage:       "start <srv-name>",
	Example:     "start srv0",
	SubCmds:     []cmd{},
	Handler: func(conn net.Conn, args []string) {
		var name string
		if len(args) > 0 {
			name = args[0]
		} else {
			fmt.Fprintf(conn, "[mmm] This command requires that you specify the name of the server to start.\n")
			return
		}

		if _, ok := instance.Instances[name]; !ok {
			fmt.Fprintf(conn, "[mmm] A server instance with the name: %s does not exist.\n", name)
			return
		}

		port := instance.Instances[name]

		for _, r := range instance.Running {
			if r.Name == name {
				fmt.Fprintf(conn, "[mmm] Server %s already running.\n", name)
				return
			}
		}

		instance.RegisterServerInstance(name, port)
		fmt.Fprintf(conn, "[mmm] Starting server %s....\n", name)
	},
}
