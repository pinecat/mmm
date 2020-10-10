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

		// Special case for quit
		if string(data) == "quit" || string(data) == "q" {
			conn.Write([]byte("[mmm] Quit command detected.  Press ENTER to return to your shell."))
			log.Trace().Msg("[mmm] Client disconnected.")
			conn.Close()
			return
		}

		// Print detailed info about client commands
		log.Trace().Msg("[mmm] (" + conn.RemoteAddr().String() + ") " + string(data))

		// Check for a command
		var isValid bool = false
		data = strings.TrimSuffix(string(data), "\n")
		split := strings.Split(data, " ")
		isValid, initCmd := cmd.Trigger(split[0], cmd.Registry)

		if !isValid {
			isValid, initCmd = cmd.SubTrigger(split[0], cmd.Registry)
		}

		if isValid {
			finCmd, args := cmd.GetDeepest(initCmd, split)
			args = args[1:]
			if len(args) >= 1 {
				if args[0] == "help" {
					finCmd.Help(conn)
				} else {
					finCmd.Handler(finCmd, conn, args)
				}
			} else {
				finCmd.Handler(finCmd, conn, args)
			}
		}

		// Check if there was a valid command, if not send an error message
		if !isValid {
			conn.Write([]byte("[mmm] Invalid command or syntax.\n"))
			log.Trace().Msg("[mmm] Client sent invalid command or syntax.")
		}
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
