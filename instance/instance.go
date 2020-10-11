package instance

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pinecat/mmm/util"
	"github.com/rs/zerolog/log"
)

func Init() {

}

func Download(version string) (bool, string, error) {
	// Get the manifest, we will need this no matter what
	manifestBytes, err := RetrieveFile(MANIFEST)
	if err != nil {
		return false, version, err
	}
	md, _ := GetManifestJSON(manifestBytes)

	// If the version passed is 'latest', use the latest version from the manifest file
	if version == "latest" {
		version = md.Latest.Release
	}

	// Search for the correct version
	for _, v := range md.Versions {
		if version == v.ID {
			versionBytes, err := RetrieveFile(v.URL)
			if err != nil {
				return false, version, err
			}
			gotVer, _ := GetVersionJSON(versionBytes)

			// Now download the actual server
			serverBytes, err := RetrieveFile(gotVer.Downloads.Server.URL)
			if err != nil {
				return false, version, err
			}

			// Now write the jar to a file
			WriteServerJar(serverBytes, version)

			// TODO: Check jar against sha1 signature

			return true, version, nil
		}
	}

	// If we loop through and can't find the version, return false, but nil
	return false, version, nil
}

func RetrieveFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Info().Msgf("[mmm] Unable to retrieve the file at: %s.", url)
		log.Info().Msg("[mmm] " + err.Error() + ".")
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info().Msg("[mmm] " + err.Error() + ".")
		return nil, err
	}

	return bytes, nil
}

func GetManifestJSON(data []byte) (ManifestData, error) {
	var md ManifestData
	err := json.Unmarshal(data, &md)
	if err != nil {
		log.Info().Msgf("[mmm] %s.", err.Error())
		return ManifestData{}, err
	}
	return md, nil
}

func GetVersionJSON(data []byte) (Version, error) {
	var v Version
	err := json.Unmarshal(data, &v)
	if err != nil {
		log.Info().Msgf("[mmm] %s.", err.Error())
		return Version{}, err
	}
	return v, nil
}

func WriteServerJar(data []byte, version string) (int64, error) {
	out, err := os.Create(util.Mmmdir + "/server" + version + ".jar")
	if err != nil {
		log.Info().Msgf("[mmm] %s.", err.Error())
		return 0, err
	}
	defer out.Close()

	reader := bytes.NewReader(data)

	n, err := io.Copy(out, reader)
	if err != nil {
		log.Info().Msgf("[mmm] %s.", err.Error())
		return 0, err
	}

	return n, nil
}

func GenEula(sdir string) {

}
