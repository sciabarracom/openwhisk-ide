package wskide

import (
	"time"

	"github.com/pkg/browser"
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

func debugParse(cmd string) bool {
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
	whiskDeploy()
	ideDeploy(dir)
	const url = "http://localhost:3000/"
	time.Sleep(2 * time.Second)
	browser.OpenURL(url)
}

// Stop openwhisk-ide
func Stop() {
	ideDestroy()
	whiskDestroy()
}

// Main entrypoint for wskide
func Main() {
	cmd := kingpin.Parse()
	if *VerboseFlag {
		log.SetLevel(log.TraceLevel)
	}
	if !(whiskParse(cmd) || ideParse(cmd) || debugParse(cmd)) {
		kingpin.Usage()
	}
}
