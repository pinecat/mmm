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
	Handler: func(conn net.Conn, args []string) {
		// This command on its own does not actually do anything
		conn.Write([]byte("[mmm] The create command must be used with a subcommand.\n"))
	},
}

var cmdCreateServer cmd = cmd{
	Name:        "server",
	Aliases:     []string{"s"},
	Description: "Create and deploy a Minecraft server instance.",
	Type:        "SubCommand",
	Usage:       "create server [version] [port] [name]",
	Example:     "cs 1.16.2",
	SubCmds:     []cmd{},
	Handler: func(conn net.Conn, args []string) {
		// Set default for version
		version := "latest"
		if len(args) > 0 {
			version = args[0]
		}

		// Set default for port
		port := ""
		if len(args) > 1 {
			port = args[1]
		}

		// Set default for name
		name := ""
		if len(args) > 2 {
			name = args[2]
		}

		// Output based on version
		if version == "latest" {
			fmt.Fprintf(conn, "[mmm] Tring to create server with the %s version.\n", version)
		} else {
			fmt.Fprintf(conn, "[mmm] Tring to create server with version %s.\n", version)
		}

		// Register the instance
		name, port, err := instance.NewServer(name, port)
		if err != nil {
			log.Info().Msgf("[mmm] %s.", err.Error())
			fmt.Fprintf(conn, "[mmm] Error registering the server.")
			return
		}
		if port == "0" {
			log.Info().Msg("[mmm] Invalid port.")
			fmt.Fprintf(conn, "[mmm] Invalid port.\n")
			return
		}

		// Actually create the instance now
		created, v, err := instance.Download(version, name)
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
			log.Trace().Msgf("[mmm] The latest Minecraft version is: %s.", v)
			fmt.Fprintf(conn, "[mmm] The latest Minecraft version is: %s.\n", v)
		}

		log.Trace().Msgf("[mmm] Sucessfully downloaded %s server jar.", v)
		log.Trace().Msgf("[mmm] Created server: %s on port: %s.", name, port)
		fmt.Fprintf(conn, "[mmm] Successfully downloaded %s server jar.\n", v)
		fmt.Fprintf(conn, "[mmm] Created server: %s on port: %s.\n", name, port)

		// Start the server
		instance.RegisterServerInstance(name, port)
		log.Trace().Msgf("[mmm] Started server: %s.", name)
		fmt.Fprintf(conn, "[mmm] Started server: %s.\n", name)
	},
}
