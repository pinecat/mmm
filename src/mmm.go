package main

// #include "mmm.h"
import "C"

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"
)

/*
	help
		Prints a help menu.
	params:
		none
	returns:
		void
*/
func help() {
	// TODO: Print a help menu
}

/*
	download
		Downloads the latest minecraft server version from Mojang.
	params:
		string - pwd - The present working directory
	returns
		err - error - If there was an error
		err - nil - If there was NOT an error
*/
func download(pwd string) error {
	// Location to download the server jar from
	//	The latest Minecraft version at this time is 1.16.3
	fileUrl := "https://launcher.mojang.com/v1/objects/f02f4473dbf152c23d7d484952121db0b36698cb/server.jar"

	// Get the file data from the above URL
	resp, err := http.Get(fileUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create file to write data to
	out, err := os.Create(pwd + "/mcserver_1.16.3.jar")
	if err != nil {
		return err
	}
	defer out.Close()

	// Write data to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func genEula(pwd string) error {
	comment := "#By changing the setting below to TRUE you are indicating your agreement to our EULA (https://account.mojang.com/documents/minecraft_eula).\n"
	dt := time.Now().String() + "\n"
	eula := "eula=true\n"

	fmt.Printf("Accept the EULA? [Y/n] ")
	ch := C.getchar()
	if ch == 'n' || ch == 'N' {
		eula = "eula=false\n"
	}

	out, err := os.Create(pwd + "/eula.txt"); if err != nil { return err }
	defer out.Close()
	_, err = out.WriteString(comment + dt + eula)

	if ch == 'n' || ch == 'N' {
		fmt.Printf("mmm: EULA not accepted.  Cannot start server.  Writing EULA and exiting....\n")
		os.Exit(0)
	}

	return nil
}

func cmdExists(cmd string) (string, bool) {
	path, err := exec.LookPath(cmd)
	return path, err == nil
}

func runInstance(pwd string) {
	// Check if java exists
	path, exists := cmdExists("java")
	if exists {
		// If java exists, fork and run the instance
		pid := C.fork()

		// Check for PID == -1 (error forking)
		if (pid == -1) {
			fmt.Printf("mmm: Unable to fork a new process.\n")
		}

		// If we are in the child process...
		if (pid != 0) {
			// Do stuff...
			cmd := &exec.Cmd {
				Path: path,
				Args: []string{path, "-jar", pwd + "/mcserver_1.16.3.jar"},
				Stdout: os.Stdout,
				Stderr: os.Stderr,
			}
			cmd.Run()
		}
	} else {
		fmt.Printf("mmm: Could not find Java on your system.  Please ensure that Java 1.8 is installed and is in your PATH.\n")
		os.Exit(1)
	}
}

func main() {
	// Get present working directory
	pwd, err := os.Getwd(); if err != nil { fmt.Printf("mmm: Permission denied.\n"); os.Exit(1) }

	// Download the file
	err = download(pwd); if err != nil { fmt.Printf("mmm: Unable to download or write server jar file.\n"); os.Exit(1) }

	err = genEula(pwd); if err != nil { fmt.Printf("mmm: Could not write eula.txt.\n"); os.Exit(1) }

	runInstance(pwd)
}