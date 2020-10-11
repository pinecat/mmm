package cmd

import (
	"net"
)

type cmdif interface {
	Help(net.Conn)
}

type cmd struct {
	Name        string
	Aliases     []string
	Description string
	Type        string
	Usage       string
	Example     string
	SubCmds     []cmd
	Handler     Action
}

type Action func(cmd, net.Conn, []string)

var Registry []cmd

func Register() {
	Registry = append(Registry,
		cmdHelp,
		cmdCreate,
		cmdRemove,
		cmdJava,
	)
}

func (c cmd) Help(conn net.Conn) {
	conn.Write([]byte("[mmm] ========== Str " + c.Name + " Help Menu ==========\n"))
	conn.Write([]byte("\tName: " + c.Name + "\n"))
	if len(c.Aliases) > 0 {
		conn.Write([]byte("\tAliases: "))
		for i, a := range c.Aliases {
			conn.Write([]byte(a))
			if i < len(c.Aliases)-1 {
				conn.Write([]byte(", "))
			} else {
				conn.Write([]byte("\n"))
			}
		}
	}
	conn.Write([]byte("\tDescription: " + c.Description + "\n"))
	conn.Write([]byte("\tType: " + c.Type + "\n"))
	conn.Write([]byte("\tUsage: " + c.Usage + "\n"))
	conn.Write([]byte("\tExample: " + c.Example + "\n"))
	if len(c.SubCmds) > 0 {
		conn.Write([]byte("\tSubcommands: "))
		for i, a := range c.SubCmds {
			conn.Write([]byte(a.Name))
			if i < len(c.Aliases)-1 {
				conn.Write([]byte(", "))
			} else {
				conn.Write([]byte("\n"))
			}
		}
	}
	conn.Write([]byte("[mmm] ========== End " + c.Name + " Help Menu ==========\n"))
}

func Trigger(cmdStr string, cmds []cmd) (bool, cmd) {
	for _, c := range cmds {
		if cmdStr == c.Name {
			return true, c
		}

		for _, a := range c.Aliases {
			if cmdStr == a {
				return true, c
			}
		}
	}
	return false, cmd{}
}

func SubTrigger(cmdStr string, cmds []cmd) (bool, cmd) {
	for _, c := range cmds {
		for _, a := range c.Aliases {
			if len(cmdStr) > 0 && string(cmdStr[0]) == a {
				if len(cmdStr) > 1 {
					return SubTrigger(cmdStr[1:], c.SubCmds)
				}
				return true, c
			}
		}
	}
	return false, cmd{}
}

func GetDeepest(prevCmd cmd, args []string) (cmd, []string) {
	if len(prevCmd.SubCmds) > 0 {
		if len(args) > 1 {
			found, c := Trigger(args[1], prevCmd.SubCmds)
			if found {
				return GetDeepest(c, args[1:])
			}
		}
	}
	return prevCmd, args
}
