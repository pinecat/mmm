package cli

import (
	"fmt"
	"net"
	"os"
)

func Create(port string, version string, sport string, name string) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Printf("mmm: Error connecting to daemon: %s", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	if name != "" {
		fmt.Fprintf(conn, "cs %s %s %s", version, sport, name)
	} else {
		fmt.Fprintf(conn, "cs %s %s", version, sport)
	}
}

func StartServ(port string, name string) {
	if name == "" {
		fmt.Printf("mmm: Please specify server name with the -name flag")
		os.Exit(1)
	}
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Printf("mmm: Error connecting to daemon: %s", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "start %s", name)
}

func Stop(port string, name string) {
	if name == "" {
		fmt.Printf("mmm: Please specify server name with the -name flag")
		os.Exit(1)
	}
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Printf("mmm: Error connecting to daemon: %s", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "stop %s", name)
}

func Remove(port string, name string) {
	if name == "" {
		fmt.Printf("mmm: Please specify server name with the -name flag")
		os.Exit(1)
	}
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Printf("mmm: Error connecting to daemon: %s", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "rs %s", name)
}
