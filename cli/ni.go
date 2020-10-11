package cli

import (
	"bufio"
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
		fmt.Fprintf(conn, "create server %s %s %s\n", version, sport, name)
	} else {
		fmt.Fprintf(conn, "create server %s %s\n", version, sport)
	}

	reader := bufio.NewReader(conn)
	for {
		text, _ := reader.ReadString('\n')
		switch text {
		case "[mmm] Invalid port.\n":
			fmt.Printf("mmm: Invalid port\n")
			os.Exit(1)
			break
		case "[mmm] Error registering the server.\n":
			fmt.Printf("mmm: Error registering the server\n")
			os.Exit(1)
			break
		case fmt.Sprintf("[mmm] Unable to download server jar, with version: %s.  Check mmm logs for more details.\n", version):
			fmt.Printf("mmm: Unable to download server jar, with version: %s, check mmm logs for more details\n", version)
			os.Exit(1)
			break
		case fmt.Sprintf("[mmm] Version: %s, is invalid.  Please specifiy a valid Minecraft version.\n", version):
			fmt.Printf("mmm: Version: %s, is invalid", version)
			os.Exit(1)
			break
		case "[mmm] Just for ni.\n":
			os.Exit(0)
			break
		}
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

	in := make(chan string)
	go func() {
		for {
			reader := bufio.NewReader(conn)
			text, _ := reader.ReadString('\n')
			in <- text
		}
	}()

	fmt.Fprintf(conn, "start %s\n", name)

	for {
		reply := <-in
		switch reply {
		case fmt.Sprintf("[mmm] A server instance with the name: %s does not exist.\n", name):
			fmt.Printf("mmm: Server instance with the name %s does not exist\n", name)
			os.Exit(1)
			break
		case fmt.Sprintf("[mmm] Server %s already running.\n", name):
			fmt.Printf("mmm: Server %s already running\n", name)
			os.Exit(1)
			break
		case fmt.Sprintf("[mmm] Starting server %s....\n", name):
			os.Exit(0)
			break
		}
	}
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

	in := make(chan string)
	go func() {
		for {
			reader := bufio.NewReader(conn)
			text, _ := reader.ReadString('\n')
			in <- text
		}
	}()

	fmt.Fprintf(conn, "stop %s\n", name)

	for {
		reply := <-in
		switch reply {
		case fmt.Sprintf("[mmm] A server instance with the name: %s does not exist.\n", name):
			fmt.Printf("mmm: Server instance with the name %s does not exist\n", name)
			os.Exit(1)
			break
		case fmt.Sprintf("[mmm] Server %s is not running.\n", name):
			fmt.Printf("mmm: Server %s is not currently running\n", name)
			os.Exit(1)
			break
		case fmt.Sprintf("[mmm] Stopping server %s.\n", name):
			os.Exit(0)
			break
		}
	}
}

func List(port string) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Printf("mmm: Error connecting to daemon: %s", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "list\n")

	reader := bufio.NewReader(conn)
	for {
		text, _ := reader.ReadString('\n')
		if text != fmt.Sprintf("[mmm] Connected to daemon mmm@%s.\n", conn.RemoteAddr().String()) &&
			text != "[mmm] You may use the 'help' command to list a short help menu.\n" &&
			text != "[mmm] Use the 'quit' ('q') command or CTRL-C to close the program.\n" &&
			text != "[mmm]λ \n" {
			fmt.Printf(text)
		}
		if text == "[mmm] ========== End List Outp ==========\n" {
			os.Exit(0)
		}
	}
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

	in := make(chan string)
	go func() {
		for {
			reader := bufio.NewReader(conn)
			text, _ := reader.ReadString('\n')
			in <- text
		}
	}()

	fmt.Fprintf(conn, "remove server %s\n", name)

	for {
		reply := <-in
		switch reply {
		case fmt.Sprintf("[mmm] Cannot remove the specified server, because a server instance with the name: %s does not exist.\n", name):
			fmt.Printf("mmm: Cammot remove server, because a server instance with the name %s does not exist\n", name)
			os.Exit(1)
			break
		case fmt.Sprintf("[mmm] Could not remove server instance: %s.\n", name):
			fmt.Printf("mmm: Could not remove server instance: %s\n", name)
			os.Exit(1)
			break
		case fmt.Sprintf("[mmm] Successfully removed server instance: %s.\n", name):
			os.Exit(0)
			break
		}
	}
}
