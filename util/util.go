package util

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

// Setup some default values, just in case the config
//	file is missing
var (
	Mmmdir string = "/usr/local/etc/mmm"
	Editor string = "vi"
	Dbglvl string = "1"
)

/*
	SetupFlags
		Setup and parse all commandline flags/arguments
	params:
		none
	returns:
		rf - *bool - Run flag
		df - *bool - Daemonize flag
		sf - *bool - Stop daemon flag
*/
func SetupFlags() (*bool, *bool, *bool, string) {
	rf := flag.Bool("r", false, "Run")
	df := flag.Bool("d", false, "Daemonize")
	sf := flag.Bool("s", false, "Stop daemon")
	pf := flag.String("p", "25564", "Daemon port")
	flag.Parse()
	return rf, df, sf, *pf
}

/*
	SetupSignals
		Setup channel and define signals to interrupt on
	params:
		none
	returns:
		ch - chan os.Signal - The channel to notify/interrupt on
*/
func SetupSignals() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	return ch
}

/*
	ExistsDir
		Checks if a file or directory exists
	params:
		path - string - Path of the file or dir to check
	returns:
		false, err - bool, error - Indicates the file or
			directory does not exist, or an error was
			encountered
		true, nil - bool, error - Indicates the file or
			directory does exist
*/
func ExistsDir(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}

/*
	CreateDir
		Attempts to create a directory with given path
	params:
		path - string - Where to make the directory
	returns:
		err - error - Indicates an error was encountered
		nil - error - Indicates the directory was created
*/
func CreateDir(path string) error {
	err := os.Mkdir(path, 0744)
	return err
}

/*
	CreateFile
		Attempts to create a file with given filepath
	params:
		path - string - Where to make the file
	returns:
		err - error - Indicates an errer has occured
		nil - error - Indicates the file was created
*/
func CreateFile(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0744)
	if err != nil {
		return err
	}
	f.Close()
	return nil
}
