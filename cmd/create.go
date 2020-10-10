package cmd

import (
	"fmt"
	"net"

	"github.com/pinecat/mmm/instance"
	"github.com/rs/zerolog/log"
)

var cmdCreate cmd = cmd{
	Name:        "create",
	Aliases:     []string{"c"},
	Description: "Send create request to mmm daemon.",
	Type:        "Command",
	Usage:       "create ...",
	Example:     "create server 1.16.2",
	SubCmds:     []cmd{cmdCreateServer},
	Handler: func(c cmd, conn net.Conn, args []string) {
		// This command on its own does not actually do anything
		conn.Write([]byte("[mmm] The create command must be used with a subcommand.\n"))
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
		//conn.Write([]byte("[mmm] Not implemented yet ¯\\_(ツ)_/¯.\n"))
		version := "latest"
		if len(args) > 0 {
			version = args[0]
		}

		if version == "latest" {
			fmt.Fprintf(conn, "[mmm] Tring to create server with the %s version.\n", version)
		} else {
			fmt.Fprintf(conn, "[mmm] Tring to create server with version %s.\n", version)
		}

		created, v, err := instance.Download(version)
		if err != nil {
			log.Info().Msgf("[mmm] %s.", err.Error())
			fmt.Fprintf(conn, "[mmm] Unable to download server jar, with version: %s.  Check mmm logs for more details.\n", version)
			return
		}

		if !created && err == nil {
			log.Trace().Msgf("[mmm] Client requested an invalid version: %s.", version)
			fmt.Fprintf(conn, "[mmm] Version: %s, is invalid.  Please specifiy a valid Minecraft version.\n", version)
			return
		}

		if version == "latest" {
			log.Trace().Msgf("[mmm] The latest version is: %s.", v)
			fmt.Fprintf(conn, "[mmm] The latest version is: %s.\n", v)
		}

		log.Trace().Msgf("[mmm] Sucessfully downloaded %s server jar.", v)
		fmt.Fprintf(conn, "[mmm] Successfully downloaded %s server jar.\n", v)
	},
}
