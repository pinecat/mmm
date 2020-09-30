package serve

import (
	"net"
	"net/textproto"
	"os"

	"github.com/rs/zerolog/log"
)

func await(conn net.Conn, tpconn *textproto.Conn) {
	conn.Write([]byte("[mmm] Connected to daemon mmm@" + conn.LocalAddr().String() + ".\n"))
	for {
		conn.Write([]byte("[mmm]λ "))
		data, err := tpconn.ReadLine()
		if err != nil {
			log.Trace().Msg("[mmm] Client disconnected.")
			conn.Close()
			return
		}
		if string(data) == "quit" {
			conn.Write([]byte("[mmm] Quit command detected.  Press ENTER to return to your shell."))
			log.Trace().Msg("[mmm] Client disconnected.")
			conn.Close()
			return
		}
		log.Trace().Msg("[mmm] (" + conn.RemoteAddr().String() + ") " + string(data))
	}
}

func Start(port string) {
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
