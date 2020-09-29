package util

import (
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ConfigDatabase struct {
	Debug int    `env:"MMMDEBUG" env-default:"1"`
	Edit  string `env:"EDITOR" env-default:"vi"`
}

var Mmmpid string = "/mmm.pid"
var Mmmlog string = "/mmm.log"

func SetEnv(ff string, rf bool, df bool, sf bool) (*os.File, ConfigDatabase) {
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
	file, err := os.Create(ff + Mmmlog)
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

func SavePid(ff string, pid int) {
	// Try to create the file
	file, err := os.Create(ff + Mmmpid)
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

func ReadPid(ff string) {
	// Check to see if the file is there
	_, err := os.Stat(ff + Mmmpid)
	if err != nil {
		log.Info().Msg("[MMM] The mmm daemon is not running.")
		os.Exit(1)
	}

	// Try to read the file
	info, err := ioutil.ReadFile(ff + Mmmpid)
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
	err = os.Remove(ff + Mmmpid)
	if err != nil {
		log.Trace().Msg("[MMM] Cannot delete pidfile: " + ff + Mmmpid + ".  Please delete it manually.")
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
