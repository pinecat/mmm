package util

import (
	"bufio"
	"bytes"
	"flag"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

// Setup some default values, just in case the config
//	file is missing
var (
	Mmmdir string = "/usr/local/srv/mmm"
	Editor string = "vi"
	Dbglvl string = "1"
	Aceula string = "0"
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
func SetupFlags() (*bool, string, string, *bool, string, string, *bool, *bool, *bool, *bool) {
	// mmm specific flags
	rf := flag.Bool("r", false, "Run")
	pf := flag.String("p", "25564", "Daemon port")

	// Server name
	nf := flag.String("n", "", "Server name (used with -c or -s flag)")

	// Flags for creating a server
	cf := flag.Bool("c", false, "Create server (may utilize other flags)")
	vf := flag.String("v", "latest", "Server version (only used with -c flag)")
	spf := flag.String("sp", "25565", "Server port (only used with -c flag)")

	// Flags for starting and stopping a server
	sf := flag.Bool("s", false, "Start server (requires the -name flag)")
	qf := flag.Bool("q", false, "Stop server (requires the -name flag)")

	// Flags for listing and removing servers
	lf := flag.Bool("l", false, "List servers (requires -name flag)")
	df := flag.Bool("d", false, "Remove server (requires -name flag)")

	flag.Parse()
	return rf, *pf, *nf, cf, *vf, *spf, sf, qf, lf, df
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
	dirs := strings.Split(path, "/")
	var fpath string = ""
	for _, d := range dirs {
		fpath = fpath + "/" + d
		if exists, _ := ExistsDir(fpath); !exists {
			err := os.Mkdir(fpath, 0775)
			if err != nil {
				return err
			}
		}
	}
	return nil
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

func CmdExists(cmd string) (string, bool) {
	path, err := exec.LookPath(cmd)
	return path, err == nil
}

func JavaVersion() string {
	path, exists := CmdExists("java")
	if !exists {
		return "[mmm] Java executable not found.\n"
	}

	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	cmd := &exec.Cmd{
		Path:   path,
		Args:   []string{path, "-version"},
		Stdout: writer,
		Stderr: writer,
	}

	cmd.Run()

	return strings.TrimSuffix(buf.String(), "\n")
}
