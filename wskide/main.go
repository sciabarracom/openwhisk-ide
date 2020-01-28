package wskide

import "gopkg.in/alecthomas/kingpin.v2"

// Main entrypoint for wskide
func Main() {
	cmd := kingpin.Parse()
	if !(whiskParse(cmd) || ideParse(cmd)) {
		kingpin.Usage()
	}
}
