package instance

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type ManifestData struct {
	Latest struct {
		Release  string `json:"release"`
		Snapshot string `json:"snapshot"`
	} `json:"latest"`
	Versions []struct {
		ID          string    `json:"id"`
		Type        string    `json:"type"`
		URL         string    `json:"url"`
		Time        time.Time `json:"time"`
		ReleaseTime time.Time `json:"releaseTime"`
	} `json:"versions"`
}

const MANIFEST string = "https://launchermeta.mojang.com/mc/game/version_manifest.json"

func Init() {

}

func Download(version string) {

}

func GetManifest() (ManifestData, error) {
	resp, err := http.Get(MANIFEST)
	if err != nil {
		log.Info().Msg("[mmm] Unable to retrieve the manifest file.  Unable determine any server.jar's download address.")
		log.Info().Msg("[mmm] " + err.Error() + ".")
		return ManifestData{}, err
	}
	defer resp.Body.Close()

	manifestBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info().Msg("[mmm] " + err.Error() + ".")
		return ManifestData{}, err
	}

	var md ManifestData
	err = json.Unmarshal(manifestBytes, &md)
	if err != nil {
		log.Info().Msg("[mmm] Unable to parse JSON from the manifest file.")
		log.Info().Msg("[mmm] " + err.Error() + ".")
		return ManifestData{}, err
	}

	return md, nil
}
