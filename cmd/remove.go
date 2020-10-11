package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/pinecat/mmm/instance"
	"github.com/pinecat/mmm/util"
	"github.com/rs/zerolog/log"
)

var cmdRemove cmd = cmd{
	Name:        "remove",
	Aliases:     []string{"r"},
	Description: "Send remove request to mmm daemon.",
	Type:        "Command",
	Usage:       "remove ...",
	Example:     "create server srv0",
	SubCmds:     []cmd{cmdRemoveServer},
	Handler: func(conn net.Conn, args []string) {
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
	Handler: func(conn net.Conn, args []string) {
		//conn.Write([]byte("[mmm] Not implemented yet ¯\\_(ツ)_/¯.\n"))
		var name string
		if len(args) > 0 {
			name = args[0]
		} else {
			fmt.Fprintf(conn, "[mmm] This command requires that you specify the name of the server to remove.\n")
			return
		}

		if _, k := instance.Instances[name]; !k {
			fmt.Fprintf(conn, "[mmm] Cannot remove the specified server, because a server instance with the name: %s does not exist.\n", name)
			return
		}

		for _, r := range instance.Running {
			if r.Name == name {
				r.Stop()
				break
			}
		}

		delete(instance.Instances, name)

		err := os.RemoveAll(util.Mmmdir + "/" + name)
		if err != nil {
			log.Info().Msgf("[mmm] %s.", err.Error())
			fmt.Fprintf(conn, "[mmm] Could not remove server instance: %s.\n", name)
			return
		}
		fmt.Fprintf(conn, "[mmm] Successfully removed server instance: %s.\n", name)
	},
}
