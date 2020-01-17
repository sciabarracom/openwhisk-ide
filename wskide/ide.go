package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	ideCmd        = kingpin.Command("ide", "IDE")
	ideDeployCmd  = ideCmd.Command("deploy", "Create IDE deployment")
	ideDestroyCmd = ideCmd.Command("destroy", "Destroy IDE deployment")
)

func ideParse(cmd string) bool {
	switch cmd {
	case ideDeployCmd.FullCommand():
		ideDeploy()
		return true
	case ideDestroyCmd.FullCommand():
		ideDestroy()
		return true
	}
	return false
}

func ideDeploy() {
	fmt.Println("Deploying IDE")
}

func ideDestroy() {
	fmt.Println("Destroying IDE")
}
