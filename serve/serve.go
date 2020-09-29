package serve

import (
	"net"
	"net/textproto"

	"github.com/rs/zerolog/log"
)

func await(conn textproto.Conn, ra net.Addr) {
	defer conn.Close()
	for {
		data, err := conn.ReadLine()
		if err != nil {
			log.Info().Msg("[MMM] Client : " + ra.String() + " disconnected.")
			conn.Close()
			break
		}
		log.Info().Msg(data)
	}
}

func Start(pf string) {
	ln, err := net.Listen("tcp", ":"+pf)
	if err != nil {
		log.Info().Msg("[MMM] Unable to listen on port: " + pf + ".  Perhaps that port is already in use.  Quitting....")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
	}

	log.Info().Msg("[MMM] Successfully started listening on port: " + pf + ".")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Info().Msg("[MMM] A client tried to connect, but there was an error.")
			log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
		}
		log.Info().Msg("[MMM] Connected to client: " + conn.RemoteAddr().String() + ".")
		tpconn := textproto.NewConn(conn)
		go await(*tpconn, conn.RemoteAddr())
	}
}
