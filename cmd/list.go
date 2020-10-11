package cmd

import (
	"fmt"
	"net"

	"github.com/pinecat/mmm/instance"
)

var cmdList cmd = cmd{
	Name:        "list",
	Aliases:     []string{"l", "ls"},
	Description: "List all server instances and all currently running servers.",
	Type:        "Command",
	Usage:       "list",
	Example:     "ls",
	SubCmds:     []cmd{cmdListInstances, cmdListRunning},
	Handler: func(conn net.Conn, args []string) {
		fmt.Fprintf(conn, "[mmm] ========== All Instances ==========\n")
		for name, port := range instance.Instances {
			fmt.Fprintf(conn, "\t%s:%s\n", name, port)
		}
		fmt.Fprintf(conn, "[mmm] ========== Crtly Running ==========\n")
		for _, r := range instance.Running {
			fmt.Fprintf(conn, "\t%s:%s\n", r.Name, r.Port)
		}
		fmt.Fprintf(conn, "[mmm] ========== End List Outp ==========\n")
	},
}

var cmdListInstances cmd = cmd{
	Name:        "instances",
	Aliases:     []string{"i"},
	Description: "List all server instances.",
	Type:        "SubCommand",
	Usage:       "list instances",
	Example:     "li",
	SubCmds:     []cmd{},
	Handler: func(conn net.Conn, args []string) {
		fmt.Fprintf(conn, "[mmm] ========== All Instances ==========\n")
		for name, port := range instance.Instances {
			fmt.Fprintf(conn, "\t%s:%s\n", name, port)
		}
		fmt.Fprintf(conn, "[mmm] ========== End List Outp ==========\n")
	},
}

var cmdListRunning cmd = cmd{
	Name:        "running",
	Aliases:     []string{"r"},
	Description: "List all currently running server instances.",
	Type:        "SubCommand",
	Usage:       "list instances",
	Example:     "li",
	SubCmds:     []cmd{},
	Handler: func(conn net.Conn, args []string) {
		fmt.Fprintf(conn, "[mmm] ========== Crtly Running ==========\n")
		for _, r := range instance.Running {
			fmt.Fprintf(conn, "\t%s:%s\n", r.Name, r.Port)
		}
		fmt.Fprintf(conn, "[mmm] ========== End List Outp ==========\n")
	},
}
