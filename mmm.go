package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/pinecat/mmm/cli"
	"github.com/pinecat/mmm/serve"
	"github.com/pinecat/mmm/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

/*
	keepAlive
		Blocking function for go routine (daemonizing)
	params:
		ch - chan os.Signal - The channel where any
			interrupts will come from
	returns:
		void
*/
func keepAlive(ch chan os.Signal) {
	for {
		sig := <-ch
		signal.Stop(ch)
		fmt.Println("Received " + sig.String() + " signal.  Quitting....")
		os.Exit(0)
	}
}

/* main: The main function */
func main() {
	// Get cmdline flags/args
	rf, pf, nf, cf, vf, spf, sf, qf, lf, df := util.SetupFlags()

	// Unfortunately, have to do some setup for logging before reading config
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123Z})

	// Read the config files, if it exists
	//	If the config files does not exist
	//	err will still be nil, we will
	//	just use hardcoded default values
	if err := util.ReadConfig(); err != nil {
		log.Info().Msg("[mmm] " + err.Error())
	}

	// See if the mmmdir exists, if not, try to create it
	if exists, _ := util.ExistsDir(util.Mmmdir); !exists {
		if err := util.CreateDir(util.Mmmdir); err != nil {
			log.Info().Msg("[mmm] " + err.Error() + ".  Quitting....")
			os.Exit(1)
		}
	}

	// Setup logging for real now
	f := util.SetupLogging()
	if f != nil && f != os.Stderr {
		defer f.Close()
	}

	// Run the actual program
	if *rf {
		ch := util.SetupSignals()
		go keepAlive(ch)
		serve.Start(pf)
	} else if *cf {
		cli.Create(pf, vf, spf, nf)
	} else if *sf {
		cli.StartServ(pf, nf)
	} else if *qf {
		cli.Stop(pf, nf)
	} else if *lf {
		cli.List(pf)
	} else if *df {
		cli.Remove(pf, nf)
	} else {
		cli.Start(pf)
	}
}
