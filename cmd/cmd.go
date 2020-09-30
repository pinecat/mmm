package cmd

import "net"

type cmdif interface {
	Help(net.Conn)
}

type cmd struct {
	Name        string
	Aliases     []string
	Description string
	Usage       string
	Example     string
	Handler     Action
}

type Action func(cmd, net.Conn, []string)

var Registry []cmd

func Register() {
	help := cmd{
		Name:        "help",
		Aliases:     []string{"h"},
		Description: "Print a general help menu.",
		Usage:       "help",
		Example:     "help",
		Handler:     helpHandler,
	}

	cs := cmd{
		Name:        "create server",
		Aliases:     []string{"cs"},
		Description: "Create and deploy a Minecraft server instance.",
		Usage:       "cs [version]",
		Example:     "cs 1.16.2",
		Handler:     csHandler,
	}

	Registry = append(Registry, help, cs)
}

func (c cmd) Help(conn net.Conn) {
	conn.Write([]byte("[mmm] ========== Str" + c.Name + " Help Menu ==========\n"))
	conn.Write([]byte("\tName: " + c.Name + "\n"))
	conn.Write([]byte("\tAliases: "))
	for i, a := range c.Aliases {
		conn.Write([]byte(a))
		if i < len(c.Aliases)-1 {
			conn.Write([]byte(", "))
		} else {
			conn.Write([]byte("\n"))
		}
	}
	conn.Write([]byte("\tDescription: " + c.Description + "\n"))
	conn.Write([]byte("\tUsage: " + c.Usage + "\n"))
	conn.Write([]byte("\tExample: " + c.Example + "\n"))
	conn.Write([]byte("[mmm] ========== End" + c.Name + " Help Menu ==========\n"))
}
