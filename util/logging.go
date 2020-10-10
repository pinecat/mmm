package util

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

/*
	SetupLogging
		Initiates logging based on the debug level
		0 - Silent, no output at all
		1 - Errors, write to stderr
		2 - Error, write to a logfile
		3 - Verbose, errors and other output, write to stderr
		4 - Verbose, errors and other output, write to logfile
	params:
		none
	returns
		void
*/
func SetupLogging() *os.File {
	var lvl zerolog.Level = zerolog.InfoLevel
	var fil *os.File = os.Stderr
	var err error

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: fil, TimeFormat: time.RFC1123Z})

	switch Dbglvl {
	case "0": // Silent
		lvl = zerolog.WarnLevel
		fil = os.Stderr
		break
	case "1": // Errors (stderr)
		lvl = zerolog.InfoLevel
		fil = os.Stderr
		break
	case "2": // Errors (logfil)
		lvl = zerolog.InfoLevel
		fil, err = openLog()
		if err != nil {
			log.Info().Msg("[mmm] " + err.Error())
		}
		break
	case "3": // Verbose (stderr)
		lvl = zerolog.TraceLevel
		fil = os.Stderr
	case "4": // Verbose (logfil)
		lvl = zerolog.TraceLevel
		fil, err = openLog()
		if err != nil {
			log.Info().Msg("[mmm] " + err.Error())
		}
		break
	}
	zerolog.SetGlobalLevel(lvl)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: fil, TimeFormat: time.RFC1123Z})
	if fil != os.Stderr {
		log.Info().Msg("")
		log.Info().Msg("")
		log.Info().Msg("====[mmm] Log Started At: " + time.Now().Format("2006-01-02 15:04:05 Monday") + " [mmm]====")
	}

	return fil
}

/*
	openLog
		Opens a log file to write logs to.  Will create the file if it does not exist,
		otherwise will append to the end of the file
	params:
		none
	returns:
		nil, err - *os.File, error - Indicates an error has occured, and the file,
			cannot be opened or created
		f, nil = *os.File, error - Returns a pointer to the log file
			WARNING: This will need to be closed eventually
*/
func openLog() (*os.File, error) {
	f, err := os.OpenFile(Mmmdir+"/mmm.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0775)
	if err != nil {
		return nil, err
	}
	return f, nil
}
