package main

import (
	"flag"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/pinecat/mmm/serve"
	"github.com/pinecat/mmm/util"
	"github.com/rs/zerolog/log"
)

func main() {
	// Setup cmdline flags and parse
	rf := flag.Bool("r", false, "Run mmm")
	df := flag.Bool("d", false, "Daemonize")
	sf := flag.Bool("s", false, "Stop daemon")
	ff := flag.String("f", "/usr/local/etc/mmm", "Config/server dir")
	pf := flag.String("p", "25564", "Port for the mmm daemon")
	_ = flag.String("mp", "25565", "Minecraft server port")
	flag.Parse()

	// Remove trailing slash from filepath if there
	if (*ff)[len(*ff)-1] == '/' {
		*ff = (*ff)[0 : len(*ff)-2]
	}

	// Set some program settings (debug level, etc)
	logfile, cfg := util.SetEnv(*ff, *rf, *df, *sf)

	// Perform an action based on the flag
	if *rf {
		// Run the daemon program normally (i.e. in the foreground)
		log.Info().Msg("[MMM] Log Level: " + strconv.Itoa(cfg.Debug) + ".")
		log.Info().Msg("[MMM] Using config & server dir: " + *ff + ".")

		// Setup signals
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

		go func() {
			sig := <-ch
			signal.Stop(ch)
			log.Info().Msg("Received " + sig.String() + " signal.  Quitting....")

			// Remove the pidfile
			err := os.Remove(*ff + util.Mmmpid)
			if err != nil {
				log.Trace().Msg("[MMM] Cannot delete pidfile: " + *ff + util.Mmmpid + ".  Please delete it manually.")
				log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
			}

			logfile.Close()
			os.Exit(0)
		}()

		serve.Start(*pf)
	} else if *df {
		// Daemonize it (will call the program again but with the -r
		//	flag to run in background)
		log.Info().Msg("[MMM] Starting the mmm daemon.")
		_, err := os.Stat(*ff + util.Mmmpid)
		if err == nil {
			log.Info().Msg("[MMM] The mmm daemon is already running.")
			log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
			os.Exit(1)
		}
		cmd := exec.Command(os.Args[0], "-r")
		cmd.Start()
		log.Info().Msg("[MMM] Started the mmm daemon [" + strconv.Itoa(cmd.Process.Pid) + "].")
		util.SavePid(*ff, cmd.Process.Pid)
		logfile.Close()
		os.Exit(0)
	} else if *sf {
		// Stop the daemon
		util.ReadPid(*ff)
		logfile.Close()
		os.Exit(0)
	} else {
		// Issue commands to the daemon, after checking to see
		//	if the daemon has been started

	}
}
