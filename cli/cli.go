package cli

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

func Start(port string) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		log.Info().Msgf("[mmm] %s.  Quitting....", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	replies := make(chan string)
	inpcmds := make(chan string)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			text, _ := reader.ReadString('\n')
			if text == "[mmm]λ \n" {
				text = strings.TrimSuffix(string(text), "\n")
			}
			replies <- text
		}
	}()

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, _ := reader.ReadString('\n')
			inpcmds <- text
		}
	}()

	for {
		select {
		case in := <-replies:
			fmt.Printf("%s", in)
			break
		case out := <-inpcmds:
			fmt.Fprintf(conn, "%s", out)
			if out == "quit\n" || out == "q\n" {
				os.Exit(0)
			}
			break
		}
	}
}
