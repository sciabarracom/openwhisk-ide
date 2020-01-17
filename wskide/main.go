package main

import "gopkg.in/alecthomas/kingpin.v2"

func main() {
	cmd := kingpin.Parse()
	if !(whiskParse(cmd) || ideParse(cmd)) {
		kingpin.Usage()
	}
}
