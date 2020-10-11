package instance

import "time"

const MANIFEST string = "https://launchermeta.mojang.com/mc/game/version_manifest.json"

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

type Version struct {
	Arguments struct {
		Game []interface{} `json:"game"`
		Jvm  []interface{} `json:"jvm"`
	} `json:"arguments"`
	AssetIndex struct {
		ID        string `json:"id"`
		Sha1      string `json:"sha1"`
		Size      int    `json:"size"`
		TotalSize int    `json:"totalSize"`
		URL       string `json:"url"`
	} `json:"assetIndex"`
	Assets    string `json:"assets"`
	Downloads struct {
		Client struct {
			Sha1 string `json:"sha1"`
			Size int    `json:"size"`
			URL  string `json:"url"`
		} `json:"client"`
		ClientMappings struct {
			Sha1 string `json:"sha1"`
			Size int    `json:"size"`
			URL  string `json:"url"`
		} `json:"client_mappings"`
		Server struct {
			Sha1 string `json:"sha1"`
			Size int    `json:"size"`
			URL  string `json:"url"`
		} `json:"server"`
		ServerMappings struct {
			Sha1 string `json:"sha1"`
			Size int    `json:"size"`
			URL  string `json:"url"`
		} `json:"server_mappings"`
	} `json:"downloads"`
	ID        string `json:"id"`
	Libraries []struct {
		Downloads struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
		} `json:"downloads"`
		Name  string `json:"name"`
		Rules []struct {
			Action string `json:"action"`
			Os     struct {
				Name string `json:"name"`
			} `json:"os"`
		} `json:"rules,omitempty"`
		Natives struct {
			Osx string `json:"osx"`
		} `json:"natives,omitempty"`
		Extract struct {
			Exclude []string `json:"exclude"`
		} `json:"extract,omitempty"`
	} `json:"libraries"`
	Logging struct {
		Client struct {
			Argument string `json:"argument"`
			File     struct {
				ID   string `json:"id"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"file"`
			Type string `json:"type"`
		} `json:"client"`
	} `json:"logging"`
	MainClass              string    `json:"mainClass"`
	MinimumLauncherVersion int       `json:"minimumLauncherVersion"`
	ReleaseTime            time.Time `json:"releaseTime"`
	Time                   time.Time `json:"time"`
	Type                   string    `json:"type"`
}

type Properties struct {
	SpawProtection                 string
	MaxTickTime                    string
	QueryPort                      string
	GeneratorSettings              string
	SyncChunkWrites                string
	ForceGamemode                  string
	AllowNether                    string
	EnforceWhitelist               string
	Gamemode                       string
	BroadcastConsoleToOps          string
	EnableQuery                    string
	PlayerIdleTimeout              string
	Difficulty                     string
	BroadcastRconToOps             string
	SpawnMonsters                  string
	OpPermissionLevel              string
	Pvp                            string
	EntityBroadcastRangePercentage string
	SnooperEnabled                 string
	LevelType                      string
	EnableStatus                   string
	Hardcore                       string
	EnableCommandBlock             string
	NetworkCompressionThreshold    string
	MaxPlayers                     string
	MaxWorldSize                   string
	ResourcePackSha1               string
	FunctionPermissionLevel        string
	RconPort                       string
	ServerPort                     string
	ServerIp                       string
	SpawnNpcs                      string
	AllowFlight                    string
	LevelName                      string
	ViewDistance                   string
	ResourcePack                   string
	SpawnAnimals                   string
	WhileList                      string
	RconPassword                   string
	GenerateStructures             string
	OnlineMode                     string
	MaxBuildHeight                 string
	LevelSeed                      string
	PreventProxyConnections        string
	UseNativeTransport             string
	EnableJmxMonitoring            string
	Motd                           string
	RateLimit                      string
	EnableRcon                     string
}
