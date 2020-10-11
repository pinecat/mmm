package cmd

import (
	"fmt"
	"net"

	"github.com/pinecat/mmm/instance"
)

var cmdStop cmd = cmd{
	Name:        "stop",
	Aliases:     []string{"p"},
	Description: "Stop a server instance.",
	Type:        "Command",
	Usage:       "stop <srv-name>",
	Example:     "stop srv0",
	SubCmds:     []cmd{},
	Handler: func(conn net.Conn, args []string) {
		var name string
		if len(args) > 0 {
			name = args[0]
		} else {
			fmt.Fprintf(conn, "[mmm] This command requires that you specify the name of the server to start.\n")
			return
		}

		if _, k := instance.Instances[name]; !k {
			fmt.Fprintf(conn, "[mmm] A server instance with the name: %s does not exist.\n", name)
			return
		}

		for _, r := range instance.Running {
			if r.Name == name {
				r.Stop()
				fmt.Fprintf(conn, "[mmm] Stopping server %s.\n", name)
				return
			}
		}

		fmt.Fprintf(conn, "[mmm] Server %s is not running.\n", name)
	},
}
