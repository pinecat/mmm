package cmd

import (
	"fmt"
	"net"

	"github.com/pinecat/mmm/util"
	"github.com/rs/zerolog/log"
)

var cmdJava cmd = cmd{
	Name:        "java",
	Aliases:     []string{"j"},
	Description: "Checks if java is installed on the system.",
	Type:        "Command",
	Usage:       "java",
	Example:     "java",
	SubCmds:     []cmd{cmdJavaVersion},
	Handler: func(conn net.Conn, args []string) {
		var path string
		var exists bool
		if path, exists = util.CmdExists("java"); !exists {
			log.Trace().Msgf("[mmm] Java executable not found.")
			fmt.Fprintf(conn, "[mmm] Java executable not found.\n")
			return
		}

		log.Trace().Msgf("[mmm] Java executable exists at: %s.", path)
		fmt.Fprintf(conn, "[mmm] Java executable exists at: %s.\n", path)
	},
}

var cmdJavaVersion cmd = cmd{
	Name:        "version",
	Aliases:     []string{"v"},
	Description: "Returns the version of Java on the system.",
	Type:        "SubCommand",
	Usage:       "java version",
	Example:     "jv",
	SubCmds:     []cmd{},
	Handler: func(conn net.Conn, args []string) {
		ver := util.JavaVersion()
		log.Trace().Msgf("[mmm] Java version info:\n %s.", ver)
		fmt.Fprintf(conn, "[mmm] Java version info:\n %s\n", ver)
	},
}
