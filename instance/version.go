package instance

import "time"

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
		Downloads2 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
		} `json:"downloads,omitempty"`
		Name  string `json:"name"`
		Rules []struct {
			Action string `json:"action"`
			Os     struct {
				Name string `json:"name"`
			} `json:"os"`
		} `json:"rules,omitempty"`
		Downloads3 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				Javadoc struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"javadoc"`
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
				Sources struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"sources"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Natives2 struct {
			Osx string `json:"osx"`
		} `json:"natives,omitempty"`
		Downloads4 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Natives3 struct {
			Linux   string `json:"linux"`
			Windows string `json:"windows"`
		} `json:"natives,omitempty"`
		Downloads5 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				Javadoc struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"javadoc"`
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
				Sources struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"sources"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Downloads6 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Natives5 struct {
			Linux   string `json:"linux"`
			Windows string `json:"windows"`
		} `json:"natives,omitempty"`
		Downloads7 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				Javadoc struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"javadoc"`
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
				Sources struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"sources"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Downloads8 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Natives7 struct {
			Linux   string `json:"linux"`
			Windows string `json:"windows"`
		} `json:"natives,omitempty"`
		Downloads9 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				Javadoc struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"javadoc"`
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
				Sources struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"sources"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Downloads10 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Natives9 struct {
			Linux   string `json:"linux"`
			Windows string `json:"windows"`
		} `json:"natives,omitempty"`
		Downloads11 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				Javadoc struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"javadoc"`
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
				Sources struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"sources"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Downloads12 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Natives10 struct {
			Linux   string `json:"linux"`
			Windows string `json:"windows"`
		} `json:"natives,omitempty"`
		Downloads13 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				Javadoc struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"javadoc"`
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
				Sources struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"sources"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Downloads14 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				Javadoc struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"javadoc"`
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
				Sources struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"sources"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Natives8 struct {
			Linux   string `json:"linux"`
			Windows string `json:"windows"`
		} `json:"natives,omitempty"`
		Downloads15 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				Javadoc struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"javadoc"`
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
				Sources struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"sources"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Downloads16 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesMacos struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-macos"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Natives11 struct {
			Linux   string `json:"linux"`
			Windows string `json:"windows"`
		} `json:"natives,omitempty"`
		Downloads17 struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				NativesLinux struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-linux"`
				NativesWindows struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-windows"`
				Sources struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"sources"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
		Extract struct {
			Exclude []string `json:"exclude"`
		} `json:"extract,omitempty"`
		Natives struct {
			Linux   string `json:"linux"`
			Windows string `json:"windows"`
		} `json:"natives,omitempty"`
		Downloads struct {
			Artifact struct {
				Path string `json:"path"`
				Sha1 string `json:"sha1"`
				Size int    `json:"size"`
				URL  string `json:"url"`
			} `json:"artifact"`
			Classifiers struct {
				Javadoc struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"javadoc"`
				NativesOsx struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"natives-osx"`
				Sources struct {
					Path string `json:"path"`
					Sha1 string `json:"sha1"`
					Size int    `json:"size"`
					URL  string `json:"url"`
				} `json:"sources"`
			} `json:"classifiers"`
		} `json:"downloads,omitempty"`
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
