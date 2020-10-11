package instance

import (
	"os"
	"os/exec"

	"github.com/pinecat/mmm/util"
	"github.com/rs/zerolog/log"
)

type ServerInstance struct {
	Name string
	Port string
	Sig  chan string
}

var Instances map[string]string
var Running []*ServerInstance

func RegisterServerInstance(name string, port string) {
	s := ServerInstance{
		Name: name,
		Port: port,
		Sig:  make(chan string),
	}

	Running = append(Running, &s)
	var running string
	for i, r := range Running {
		if i == len(Running)-1 {
			running += r.Name
		} else {
			running += r.Name + ", "
		}
	}
	log.Trace().Msgf("[mmm] %s.", running)
	s.Start()
}

func (s ServerInstance) Start() {
	var path string
	var exists bool
	if path, exists = util.CmdExists("java"); !exists {
		return
	}

	jar := util.Mmmdir + "/" + s.Name + "/server.jar"
	cmd := &exec.Cmd{
		Path:   path,
		Args:   []string{path, "-jar", jar},
		Stdout: log.Logger,
		Stderr: log.Logger,
		Dir:    util.Mmmdir + "/" + s.Name,
	}

	go func() {
		go cmd.Run()

		for {
			sig := <-s.Sig
			if sig == "quit" || sig == "kill" {
				var i int
				for i = 0; i < len(Running); i++ {
					if Running[i].Name == s.Name {
						break
					}
				}
				Running[i] = Running[len(Running)-1]
				Running[len(Running)-1] = nil
				Running = Running[:len(Running)-1]
				log.Trace().Msgf("[mmm] Running: %s.", Running)

				if sig == "quit" {
					cmd.Process.Signal(os.Interrupt)
				} else if sig == "kill" {
					cmd.Process.Kill()
				}

				return
			}
		}
	}()
}

func (s ServerInstance) Stop(sig string) {
	s.Sig <- sig
}
