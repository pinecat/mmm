package util

import (
	"bufio"
	"os"
	"strings"
)

const CONFIGFILE string = "/usr/local/etc/mmm/mmm.conf"

/*
	ReadConfig
		Checks to see if /usr/local/etc/mmm/mmm.conf exists, and if it
		does, attempts to read in configuration from it
	params:
		none
	returns:
		err - error - Indicates an error occurred when
			attempting to open the file
		nil - error - Indicates the file does not exist
			or the config was read in successfully
*/
func ReadConfig() error {
	// Check if the .mmmrc file exists
	fp := CONFIGFILE
	if exists, _ := ExistsDir(fp); !exists {
		return nil
	}

	// Open the file for reading
	f, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer f.Close()

	// Create a new bufio reader
	reader := bufio.NewScanner(f)

	// Read in values and parse
	for reader.Scan() {
		key := strings.Split(reader.Text(), "=")[0]
		val := strings.Split(reader.Text(), "=")[1]

		switch strings.ToLower(key) {
		case "mmmdir":
			if val[len(val)-1] == '/' {
				val = val[0 : len(val)-1]
			}
			Mmmdir = os.ExpandEnv(val)
			break
		case "editor":
			Editor = os.ExpandEnv(val)
			break
		case "dbglvl":
			Dbglvl = val
			break
		case "aceula":
			Aceula = val
		}
	}

	return nil
}
