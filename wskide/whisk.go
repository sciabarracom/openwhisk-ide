package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	whiskCmd        = kingpin.Command("whisk", "Whisk")
	whiskDeployCmd  = whiskCmd.Command("deploy", "Create Whisk deployment")
	whiskDestroyCmd = whiskCmd.Command("destroy", "Destroy Whisk deployment")
)

func whiskParse(cmd string) bool {
	switch cmd {
	case whiskDeployCmd.FullCommand():
		whiskDeploy()
		return true
	case whiskDestroyCmd.FullCommand():
		whiskDestroy()
		return true
	}
	return false
}

func whiskDeploy() {
	fmt.Println("Deploying Whisk")
}

func whiskDestroy() {
	fmt.Println("Destroying Whisk")
}
