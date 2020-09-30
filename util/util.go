package util

import (
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ConfigDatabase struct {
	Debug int    `env:"MMMDEBUG" env-default:"5"`
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
	var level zerolog.Level

	// Set log level
	switch cfg.Debug {
	case 0: // Silent
		level = zerolog.WarnLevel
		break
	case 1: // Info
		level = zerolog.InfoLevel
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})
		break
	case 2: // Full errors
		level = zerolog.TraceLevel
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})
		break
	case 3: // Info & out to a log file
		level = zerolog.InfoLevel
		if rf || df || sf {
			logfile = saveLog(ff)
			if logfile != nil {
				writers = append(writers, zerolog.ConsoleWriter{Out: logfile, TimeFormat: time.RFC1123Z})
			}
		}
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})
		break
	case 4: // Full errors & Out to a log file
		level = zerolog.TraceLevel
		if rf || df || sf {
			logfile = saveLog(ff)
			if logfile != nil {
				writers = append(writers, zerolog.ConsoleWriter{Out: logfile, TimeFormat: time.RFC1123Z})
			}
		}
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})
		break
	case 5: // Just out to a log file
		level = zerolog.InfoLevel
		if rf || df || sf {
			logfile = saveLog(ff)
			if logfile != nil {
				writers = append(writers, zerolog.ConsoleWriter{Out: logfile, TimeFormat: time.RFC1123Z})
			}
		}
	default: // Default to Info
		level = zerolog.InfoLevel
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})
		break
	}

	mw := io.MultiWriter(writers...)
	log.Logger = zerolog.New(mw).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(level)

	return logfile, cfg
}

func saveLog(ff string) *os.File {
	// Try to create the file
	file, err := os.Create(ff + Mmmlog)
	if err != nil {
		log.Info().Msg("[MMM] Unable to open logfile for daemon.  Will only write to stderr.")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
	}

	if file == nil {
		log.Info().Msg("File is nil")
	}

	err = syscall.Access(ff+Mmmlog, syscall.O_RDWR)
	if err != nil {
		log.Info().Msg("[MMM] Logfile not writable.  Quitting....")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
		os.Exit(1)
	}

	// Flush write to disk
	file.Sync()

	// Get the writer
	return file
}

func SavePid(ff string, pid int) {
	// Try to open the file
	file, err := os.Create(ff + Mmmpid)
	if err != nil {
		log.Info().Msg("[MMM] Unable to open pidfile for daemon.  Quitting....")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
		os.Exit(1)
	}
	defer file.Close()

	err = syscall.Access(ff+Mmmpid, syscall.O_RDWR)
	if err != nil {
		log.Info().Msg("[MMM] Pidfile not writable.  Quitting....")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
		os.Exit(1)
	}

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

	err = syscall.Access(ff+Mmmpid, syscall.O_RDWR)
	if err != nil {
		log.Info().Msg("[MMM] Pidfile not readable.  Quitting....")
		log.Trace().Msg("[ERR] Trace: " + err.Error() + ".")
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
		log.Info().Msg("[MMM] Unable to read PID.  Quitting....")
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
		log.Trace().Msg("[MMM] Cannot delete pidfile: " + ff + Mmmpid + ", or pidfile does not exist.")
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
