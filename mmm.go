package main

import (
	"flag"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pinecat/mmm/serve"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ConfigDatabase struct {
	Debug int    `env:"MMMDEBUG" env-default:"1"`
	Edit  string `env:"EDITOR" env-default:"vi"`
}

var mmmpid string = "/mmm.pid"
var mmmlog string = "/mmm.log"

func setEnv(ff string, rf bool, df bool, sf bool) (*os.File, ConfigDatabase) {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})

	var cfg ConfigDatabase
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Info().Msg("[MMM] Unable to read environment variables, continuing with default.")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
	}

	var writers []io.Writer
	var logfile *os.File

	// Set log level
	switch cfg.Debug {
	case 0: // Silent
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		break
	case 1: // Info
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})
		break
	case 2: // Full errors
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})
		break
	case 3: // Info & out to a log file
		if rf || df || sf {
			logfile = saveLog(ff)
			writers = append(writers, zerolog.ConsoleWriter{Out: logfile, TimeFormat: time.RFC1123Z})
		}
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})
		break
	case 4: // Full errors & Out to a log file
		if rf || df || sf {
			logfile = saveLog(ff)
			writers = append(writers, zerolog.ConsoleWriter{Out: logfile, TimeFormat: time.RFC1123Z})
		}
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})
		break
	default: // Default to Info
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})
		break
	}

	mw := io.MultiWriter(writers...)
	log.Logger = zerolog.New(mw).With().Timestamp().Logger()

	return logfile, cfg
}

func saveLog(ff string) *os.File {
	// Try to create the file
	file, err := os.Create(ff + mmmlog)
	if err != nil {
		log.Info().Msg("[MMM] Unable to create logfile for daemon.  Quitting....")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
		os.Exit(1)
	}

	// Flush write to disk
	file.Sync()

	// Get the writer
	return file
}

func savePid(ff string, pid int) {
	// Try to create the file
	file, err := os.Create(ff + mmmpid)
	if err != nil {
		log.Info().Msg("[MMM] Unable to create pidfile for daemon.  Quitting....")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
		os.Exit(1)
	}
	defer file.Close()

	// Write PID to the pidfile
	_, err = file.WriteString(strconv.Itoa(pid))
	if err != nil {
		log.Info().Msg("[MMM] Unable to write PID to the pidfile.  Quitting....")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
		os.Exit(1)
	}

	// Flush write to disk
	file.Sync()
}

func readPid(ff string) {
	// Check to see if the file is there
	_, err := os.Stat(ff + mmmpid)
	if err != nil {
		log.Info().Msg("[MMM] The mmm daemon is not running.")
		os.Exit(1)
	}

	// Try to read the file
	info, err := ioutil.ReadFile(ff + mmmpid)
	if err != nil {
		log.Info().Msg("[MMM] Unable to read pidfile.  Quitting....")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
		os.Exit(1)
	}

	// Get the PID
	pid, err := strconv.Atoi(string(info))
	if err != nil {
		log.Info().Msg("[MMM] Pidfile corrupted.  Try deleting it and stopping the daemon manually.  Quitting....")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
		os.Exit(1)
	}

	// Find the process from the PID
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Info().Msg("[MMM] Cannot find mmm process with PID: " + strconv.Itoa(pid) + ".  Are you sure it's running?  Quitting....")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
		os.Exit(1)
	}

	// Remove the pidfile
	err = os.Remove(ff + mmmpid)
	if err != nil {
		log.Trace().Msg("[MMM] Cannot delete pidfile: " + ff + mmmpid + ".  Please delete it manually.")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
	}

	// Kill the process
	err = process.Kill()
	if err != nil {
		log.Info().Msg("[MMM] Unable to kill the mmm daemon.  Quitting....")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
		os.Exit(1)
	}

	log.Info().Msg("[MMM] Successfully killed the mmm daemon.")
}

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
	logfile, cfg := setEnv(*ff, *rf, *df, *sf)

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
			err := os.Remove(*ff + mmmpid)
			if err != nil {
				log.Trace().Msg("[MMM] Cannot delete pidfile: " + *ff + mmmpid + ".  Please delete it manually.")
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
		_, err := os.Stat(*ff + mmmpid)
		if err == nil {
			log.Info().Msg("[MMM] The mmm daemon is already running.")
			log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
			os.Exit(1)
		}
		cmd := exec.Command(os.Args[0], "-r")
		cmd.Start()
		log.Info().Msg("[MMM] Started the mmm daemon [" + strconv.Itoa(cmd.Process.Pid) + "].")
		savePid(*ff, cmd.Process.Pid)
		logfile.Close()
		os.Exit(0)
	} else if *sf {
		// Stop the daemon
		readPid(*ff)
		logfile.Close()
		os.Exit(0)
	} else {
		// Issue commands to the daemon, after checking to see
		//	if the daemon has been started

	}
}
