package instance

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/pinecat/mmm/util"
	"github.com/rs/zerolog/log"
)

func GetServers() {
	for k := range Instances {
		delete(Instances, k)
	}

	files, err := ioutil.ReadDir(util.Mmmdir)
	if err != nil {
		log.Info().Msgf("[mmm] %s.  Quitting....", err.Error())
		os.Exit(1)
	}

	for _, f := range files {
		if f.IsDir() {
			port, _ := ReadProperties(f.Name())
			Instances[f.Name()] = port
		}
	}
}

func NewServer(name string, port string) (string, string, error) {
	// Adjust port
	port = GenPort(port)
	if port == "0" {
		return name, port, nil
	}

	// If no name is supplied, use a default, generic name
	if name == "" {
		for i := 0; true; i++ {
			name = "srv" + strconv.Itoa(i)
			if _, s := Instances[name]; !s {
				Instances[name] = port
				util.CreateDir(util.Mmmdir + "/" + name)
				GenProperties(name, port)
				GenEula(name)
				return name, port, nil
			}
		}
	} else { // Otherwise try to make server with specified name
		if _, s := Instances[name]; !s {
			Instances[name] = port
			util.CreateDir(util.Mmmdir + "/" + name)
			GenProperties(name, port)
			GenEula(name)
			return name, port, nil
		} else { // If a server with that name already exists, call NewServer without supplying a name for auto generated name
			return NewServer("", port)
		}
	}

	// Will never make it here, but need to put this so the compiler doesn't freak out
	return name, port, nil
}

func RmServer() {

}

func GenPort(port string) string {
	var p int
	var err error

	if port == "" {
		p = 25565
	} else {
		p, err = strconv.Atoi(port)
		if err != nil {
			log.Info().Msgf("[mmm] %s.", err.Error())
			return "0"
		}
	}

	for _, v := range Instances {
		vport, _ := strconv.Atoi(v)
		if vport >= p {
			p = vport + 1
		}
	}

	return strconv.Itoa(p)
}

func ReadProperties(subdir string) (string, error) {
	var port string

	fp := util.Mmmdir + "/" + subdir + "/server.properties"
	if exists, _ := util.ExistsDir(fp); !exists {
		return "unknown", nil
	}

	f, err := os.Open(fp)
	if err != nil {
		return "unknown", err
	}
	defer f.Close()

	reader := bufio.NewScanner(f)

	for reader.Scan() {
		var key string
		var val string
		split := strings.Split(reader.Text(), "=")

		key = split[0]
		if len(split) > 1 {
			val = strings.Split(reader.Text(), "=")[1]
		}

		switch strings.ToLower(key) {
		case "server-port":
			port = val
			break
		}
	}

	return port, nil
}
