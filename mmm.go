package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/pinecat/mmm/serve"
	"github.com/pinecat/mmm/util"
)

func keepAlive(ch chan os.Signal) {
	for {
		sig := <-ch
		signal.Stop(ch)
		fmt.Println("Received " + sig.String() + " signal.  Quitting....")
		os.Remove("./mmm.pid")
		os.Exit(0)
	}
}

func main() {
	rf, df, sf, _ := util.SetupFlags()

	if *rf {
		ch := util.SetupSignals()
		go keepAlive(ch)
		serve.Start()
	}

	if *df {
		cmd := exec.Command(os.Args[0], "-r")
		cmd.Start()
		os.Create("./mmm.pid")
		if err := ioutil.WriteFile("./mmm.pid", []byte(strconv.Itoa(cmd.Process.Pid)), 0644); err != nil {
			fmt.Printf("%v", err)
		}
		fmt.Printf("PID of the child process: %d.\n", cmd.Process.Pid)
		os.Exit(0)
	}

	if *sf {
		pidstr, _ := ioutil.ReadFile("./mmm.pid")
		pid, _ := strconv.Atoi(string(pidstr))
		syscall.Kill(pid, syscall.SIGKILL)
	}
}
