package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"syscall"

	"github.com/rs/zerolog/log"
)

func Daemonize(prg string, port string) {
	// See if there is already a pidfile
	// 	This would indicate the daemon is
	//	already running
	if exists, err := ExistsDir(Mmmdir + "/mmm.pid"); exists || err == nil {
		log.Info().Msg("[mmm] Daemon already running.")
		os.Exit(1)
	}

	// Try to create the pidfile.  If it fails
	//	then just quit, since we don't want to
	//	start the daemon without a pidfile
	if err := CreateFile(Mmmdir + "/mmm.pid"); err != nil {
		log.Info().Msg("[mmm] " + err.Error() + ".  Quitting....")
		os.Exit(1)
	}

	// Create the command and run it
	cmd := exec.Command(os.Args[0], "-p="+port, "-r")
	cmd.Start()

	// We shouldn't need to check if we can read/write to this, since we checked i
	//	above before running the command
	if err := ioutil.WriteFile(Mmmdir+"/mmm.pid", []byte(strconv.Itoa(cmd.Process.Pid)), 0744); err != nil {
		fmt.Printf("%v", err)
	}

	// Print some verbose info
	log.Trace().Msg("[mmm] Daemon running with PID: " + strconv.Itoa(cmd.Process.Pid) + ".")
}

/*
	StopDaemon
		Stops the current running mmm daemon,
		if there is one

*/
func StopDaemon() {
	// Read in PID from the pidfile
	pidstr, err := ioutil.ReadFile(Mmmdir + "/mmm.pid")
	if err != nil {
		log.Info().Msg("[mmm] " + err.Error() + ".  Quitting....")
		os.Exit(1)
	}

	// Remove the PID file.
	os.Remove(Mmmdir + "/mmm.pid")

	// Convert the PID to an integer
	pid, _ := strconv.Atoi(string(pidstr))

	// Send SIGKILL the mmm daemon
	err = syscall.Kill(pid, syscall.SIGKILL)
	if err != nil {
		log.Info().Msg("[mmm] " + err.Error() + ".  Quitting....")
		os.Exit(1)
	}

	// Print some verbose info
	log.Trace().Msg("[mmm] Killed mmm daemon with PID: " + string(pidstr) + ".")
}
