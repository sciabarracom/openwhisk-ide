package wskide

import "gopkg.in/alecthomas/kingpin.v2"
import log "github.com/sirupsen/logrus"

// VerboseFlag is flag for verbose
var VerboseFlag = kingpin.Flag("verbose", "Verbose").
	Short('v').Default("false").Bool()

// Main entrypoint for wskide
func Main() {
	cmd := kingpin.Parse()
	if *VerboseFlag {
		log.SetLevel(log.TraceLevel)
	}
	if !(whiskParse(cmd) || ideParse(cmd)) {
		kingpin.Usage()
	}
}
