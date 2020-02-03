package wskide

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/mitchellh/go-homedir"
)

func compareDockerVersion() error {
	vA := semver.New(MinDockerVersion)
	vB := semver.New(strings.TrimSpace(dockerVersion()))
	if vB.Compare(*vA) == -1 {
		fmt.Printf("Installed docker version %s is no longer supported", vB)
		return errors.New("Error")
	}
	return nil
}

func inHomePath(dir string) error {
	homePath, err := homedir.Dir()
	if err != nil {
		return err
	}
	dir, err = filepath.Abs(dir)
	if err != nil {
		return err
	}
	if !strings.HasPrefix(dir, homePath) {
		return fmt.Errorf("Path %s is not subdir of homeDir %s", dir, homePath)
	}
	return nil
}
