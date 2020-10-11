package instance

import (
	"io/ioutil"
	"os"

	"github.com/pinecat/mmm/util"
	"github.com/rs/zerolog/log"
)

func GetServers() {
	files, err := ioutil.ReadDir(util.Mmmdir)
	if err != nil {
		log.Info().Msgf("[mmm] %s.  Quitting....", err.Error())
		os.Exit(1)
	}

	for _, f := range files {
		if f.IsDir() {
			log.Info().Msgf("Dir: %s.", f.Name())
		}
	}
}
