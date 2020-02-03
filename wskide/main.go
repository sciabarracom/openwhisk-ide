package wskide

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

// VerboseFlag is flag for verbose
var VerboseFlag = kingpin.Flag("verbose", "Verbose").
	Short('v').Default("false").Bool()

var (
	startCmd    = kingpin.Command("start", "Create deployment")
	startDirArg = startCmd.Arg("dir", "Project dir").Required().String()
	stopCmd     = kingpin.Command("stop", "Destroy deployment")
)

func cmdParse(cmd string) bool {
	switch cmd {
	case startCmd.FullCommand():
		Start(*startDirArg)
		return true
	case stopCmd.FullCommand():
		Stop()
		return true
	}
	return false
}

// Start openwhisk-ide
func Start(dir string) {
	var err error
	err = compareDockerVersion()
	FatalIf(err)
	err = inHomePath(dir)
	FatalIf(err)
	dockerRunOpenWhisk()
	dockerRunIde()
}

// Stop openwhisk-ide
func Stop() {
	dockerRmIde()
	dockerRmOpenWhisk()
}

// Main entrypoint for wskide
func Main() {
	cmd := kingpin.Parse()
	if *VerboseFlag {
		log.SetLevel(log.TraceLevel)
	}
	if !(whiskParse(cmd) || ideParse(cmd) || cmdParse(cmd)) {
		kingpin.Usage()
	}
}
