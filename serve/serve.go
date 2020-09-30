package serve

import (
	"net"
	"net/textproto"
	"os"
	"strings"

	"github.com/pinecat/mmm/cmd"
	"github.com/rs/zerolog/log"
)

func await(conn net.Conn, tpconn *textproto.Conn) {
	// Print some connection info
	conn.Write([]byte("[mmm] Connected to daemon mmm@" + conn.LocalAddr().String() + ".\n"))
	conn.Write([]byte("[mmm] You may use the 'help' command to list a short help menu.\n"))
	conn.Write([]byte("[mmm] Use the 'quit' ('q') command or CTRL-C to close the program.\n"))
	for {
		// Send a PS to the client
		conn.Write([]byte("[mmm]λ "))

		// Read in from the connection per line
		data, err := tpconn.ReadLine()

		// If there is an error, disconnect the client
		// 	They most likely just disconnected themselves
		if err != nil {
			log.Trace().Msg("[mmm] Client disconnected.")
			conn.Close()
			return
		}

		// Check for a command
		var isValid bool = false
		data = strings.TrimSuffix(string(data), "\n")
		args := strings.Split(data, " ")
		for _, cmd := range cmd.Registry {
			if args[0] == cmd.Name {
				isValid = true
				cmd.Handler(cmd, conn, args)
				break
			} else if len(args) > 1 && args[0]+" "+args[1] == cmd.Name {
				isValid = true
				cmd.Handler(cmd, conn, args)
				break
			} else {
				for _, a := range cmd.Aliases {
					if args[0] == a {
						isValid = true
						cmd.Handler(cmd, conn, args)
						break
					}
				}
			}
		}

		// Special case for quit
		if string(data) == "quit" || string(data) == "q" {
			conn.Write([]byte("[mmm] Quit command detected.  Press ENTER to return to your shell."))
			log.Trace().Msg("[mmm] Client disconnected.")
			conn.Close()
			return
		}

		// Check if there was a valid command, if not send an error message
		if !isValid {
			conn.Write([]byte("[mmm] Invalid command or syntax.\n"))
		}

		// Print detailed info about client commands
		log.Trace().Msg("[mmm] (" + conn.RemoteAddr().String() + ") " + string(data))
	}
}

func Start(port string) {
	cmd.Register()
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Info().Msg("[mmm] " + err.Error() + ".  Quitting....")
		os.Exit(1)
	}
	log.Trace().Msg("[mmm] Started listening on port: " + port + ".")
	for {
		conn, _ := ln.Accept()
		log.Trace().Msg("[mmm] Got a client.")
		go await(conn, textproto.NewConn(conn))
	}
}
