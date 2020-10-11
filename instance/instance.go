package instance

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/pinecat/mmm/util"
	"github.com/rs/zerolog/log"
)

func Init() {
	Instances = make(map[string]string)
	Running = make([]*ServerInstance, 0)
	GetServers()

	for name, port := range Instances {
		RegisterServerInstance(name, port)
	}
}

func Download(version string, subpath string) (bool, string, error) {
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
			WriteServerJar(serverBytes, version, subpath)

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

func WriteServerJar(data []byte, version string, subpath string) (int64, error) {
	out, err := os.Create(util.Mmmdir + "/" + subpath + "/server.jar")
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

func GenEula(name string) error {
	comment := "#By changing the setting below to TRUE you are indicating your agreement to our EULA (https://account.mojang.com/documents/minecraft_eula).\n"
	dt := fmt.Sprintf("#%s\n", time.Now().Format("Mon Jan 02 15:04:05 MST 2006"))
	eula := "eula=true\n"

	out, err := os.Create(util.Mmmdir + "/" + name + "/eula.txt")
	if err != nil {
		log.Info().Msgf("[mmm] %s.", err.Error())
		return err
	}
	defer out.Close()
	_, err = out.WriteString(comment + dt + eula)
	if err != nil {
		log.Info().Msgf("[mmm] %s.", err.Error())
		return err
	}

	return nil
}

func GenProperties(name string, port string) error {
	f, err := os.OpenFile(util.Mmmdir+"/"+name+"/server.properties", os.O_CREATE|os.O_RDWR, 0775)
	if err != nil {
		log.Info().Msgf("[mmm] %s.", err.Error())
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)

	prop := Properties{
		SpawProtection:                 "16",
		MaxTickTime:                    "60000",
		GeneratorSettings:              "",
		SyncChunkWrites:                "true",
		ForceGamemode:                  "false",
		AllowNether:                    "true",
		EnforceWhitelist:               "false",
		Gamemode:                       "survival",
		BroadcastConsoleToOps:          "true",
		EnableQuery:                    "false",
		PlayerIdleTimeout:              "0",
		Difficulty:                     "easy",
		BroadcastRconToOps:             "true",
		SpawnMonsters:                  "true",
		OpPermissionLevel:              "4",
		Pvp:                            "true",
		EntityBroadcastRangePercentage: "100",
		SnooperEnabled:                 "true",
		LevelType:                      "default",
		EnableStatus:                   "true",
		Hardcore:                       "false",
		EnableCommandBlock:             "false",
		NetworkCompressionThreshold:    "256",
		MaxPlayers:                     "20",
		MaxWorldSize:                   "29999984",
		ResourcePackSha1:               "",
		FunctionPermissionLevel:        "2",
		RconPort:                       "25575",
		ServerPort:                     port,
		ServerIp:                       "",
		SpawnNpcs:                      "true",
		AllowFlight:                    "false",
		LevelName:                      "world",
		ViewDistance:                   "10",
		ResourcePack:                   "",
		SpawnAnimals:                   "true",
		WhileList:                      "false",
		RconPassword:                   "",
		GenerateStructures:             "true",
		OnlineMode:                     "true",
		MaxBuildHeight:                 "256",
		LevelSeed:                      "",
		PreventProxyConnections:        "false",
		UseNativeTransport:             "true",
		EnableJmxMonitoring:            "false",
		Motd:                           "A Minecraft Server",
		RateLimit:                      "0",
		EnableRcon:                     "false",
	}

	writer.WriteString("#Minecraft server properties\n")
	writer.WriteString(fmt.Sprintf("#%s\n", time.Now().Format("Mon Jan 02 15:04:05 MST 2006")))

	writer.WriteString(fmt.Sprintf("spawn-protection=%s\n", prop.SpawProtection))
	writer.WriteString(fmt.Sprintf("max-tick-time=%s\n", prop.MaxTickTime))
	writer.WriteString(fmt.Sprintf("query.port=%s\n", prop.QueryPort))
	writer.WriteString(fmt.Sprintf("generator-settings=%s\n", prop.GeneratorSettings))
	writer.WriteString(fmt.Sprintf("sync-chunk-writes=%s\n", prop.SyncChunkWrites))
	writer.WriteString(fmt.Sprintf("force-gamemode=%s\n", prop.ForceGamemode))
	writer.WriteString(fmt.Sprintf("allow-nether=%s\n", prop.AllowNether))
	writer.WriteString(fmt.Sprintf("enforce-whitelist=%s\n", prop.EnforceWhitelist))
	writer.WriteString(fmt.Sprintf("gamemode=%s\n", prop.Gamemode))
	writer.WriteString(fmt.Sprintf("broadcast-console-to-ops=%s\n", prop.BroadcastConsoleToOps))
	writer.WriteString(fmt.Sprintf("enable-query=%s\n", prop.EnableQuery))
	writer.WriteString(fmt.Sprintf("player-idle-timeout=%s\n", prop.PlayerIdleTimeout))
	writer.WriteString(fmt.Sprintf("difficulty=%s\n", prop.Difficulty))
	writer.WriteString(fmt.Sprintf("broadcast-rcon-to-ops=%s\n", prop.BroadcastRconToOps))
	writer.WriteString(fmt.Sprintf("spawn-monsters=%s\n", prop.SpawnMonsters))
	writer.WriteString(fmt.Sprintf("op-permission-level=%s\n", prop.OpPermissionLevel))
	writer.WriteString(fmt.Sprintf("pvp=%s\n", prop.Pvp))
	writer.WriteString(fmt.Sprintf("entity-broadcast-range-percentage=%s\n", prop.EntityBroadcastRangePercentage))
	writer.WriteString(fmt.Sprintf("snooper-enabled=%s\n", prop.SnooperEnabled))
	writer.WriteString(fmt.Sprintf("level-type=%s\n", prop.LevelType))
	writer.WriteString(fmt.Sprintf("enable-status=%s\n", prop.EnableStatus))
	writer.WriteString(fmt.Sprintf("hardcore=%s\n", prop.Hardcore))
	writer.WriteString(fmt.Sprintf("enable-command-block=%s\n", prop.EnableCommandBlock))
	writer.WriteString(fmt.Sprintf("network-compression-threshold=%s\n", prop.NetworkCompressionThreshold))
	writer.WriteString(fmt.Sprintf("max-players=%s\n", prop.MaxPlayers))
	writer.WriteString(fmt.Sprintf("max-world-size=%s\n", prop.MaxWorldSize))
	writer.WriteString(fmt.Sprintf("resource-pack-sha1=%s\n", prop.ResourcePackSha1))
	writer.WriteString(fmt.Sprintf("function-permission-level=%s\n", prop.FunctionPermissionLevel))
	writer.WriteString(fmt.Sprintf("rcon.port=%s\n", prop.RconPort))
	writer.WriteString(fmt.Sprintf("server-port=%s\n", prop.ServerPort))
	writer.WriteString(fmt.Sprintf("server-ip=%s\n", prop.ServerIp))
	writer.WriteString(fmt.Sprintf("spawn-npcs=%s\n", prop.SpawnNpcs))
	writer.WriteString(fmt.Sprintf("allow-flight=%s\n", prop.AllowFlight))
	writer.WriteString(fmt.Sprintf("level-name=%s\n", prop.LevelName))
	writer.WriteString(fmt.Sprintf("view-distance=%s\n", prop.ViewDistance))
	writer.WriteString(fmt.Sprintf("resource-pack=%s\n", prop.ResourcePack))
	writer.WriteString(fmt.Sprintf("spawn-animals=%s\n", prop.SpawnAnimals))
	writer.WriteString(fmt.Sprintf("white-list=%s\n", prop.WhileList))
	writer.WriteString(fmt.Sprintf("rcon.password=%s\n", prop.RconPassword))
	writer.WriteString(fmt.Sprintf("generate-structures=%s\n", prop.GenerateStructures))
	writer.WriteString(fmt.Sprintf("online-mode=%s\n", prop.OnlineMode))
	writer.WriteString(fmt.Sprintf("max-build-height=%s\n", prop.MaxBuildHeight))
	writer.WriteString(fmt.Sprintf("level-seed=%s\n", prop.LevelSeed))
	writer.WriteString(fmt.Sprintf("prevent-proxy-connection=%s\n", prop.PreventProxyConnections))
	writer.WriteString(fmt.Sprintf("use-native-transport=%s\n", prop.UseNativeTransport))
	writer.WriteString(fmt.Sprintf("enable-jmx-monitoring=%s\n", prop.EnableJmxMonitoring))
	writer.WriteString(fmt.Sprintf("motd=%s\n", prop.Motd))
	writer.WriteString(fmt.Sprintf("rate-limit=%s\n", prop.RateLimit))
	writer.WriteString(fmt.Sprintf("enable-rcon=%s\n", prop.EnableRcon))

	writer.Flush()

	return nil
}
