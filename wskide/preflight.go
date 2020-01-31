package wskide

import (
	"errors"
	"fmt"
	"strings"

	"github.com/coreos/go-semver/semver"
)

func compareDockerVersion() error {
	vA := semver.New(MinDockerVersion)
	vB := semver.New(strings.TrimSpace(dockerVersion()))
	if vB.Compare(*vA) == -1 {
		fmt.Printf("La versione di docker %s installata non è piu' supportata", vB)
		return errors.New("Error")
	}
	fmt.Printf("La versione di docker installata %s è supportata", vB)
	return nil
}
