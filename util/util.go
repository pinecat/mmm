package util

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
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
	ff := flag.String("f", "/usr/local/etc/mmm", "Default dir for mmm operations")
	flag.Parse()
	return rf, df, sf, *ff
}

func SetupSignals() chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	return ch
}
