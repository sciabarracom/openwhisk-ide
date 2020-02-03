package wskide

import (
	"fmt"
	"os"
)

func ExampleCompareDockerVersion() {
	*DryRunFlag = false
	compareDockerVersion()
	// Output:
}
func ExamplePassDockerVersion() {
	// *DryRunFlag = false
	DryRunPush("19.03.5")
	compareDockerVersion()
	// Output:
	// docker version --format {{.Server.Version}}
}
func ExampleFailDockerVersion() {
	// *DryRunFlag = false
	DryRunPush("10.03.5")
	compareDockerVersion()
	// Output:
	// docker version --format {{.Server.Version}}
	// Installed docker version 10.3.5 is no longer supported
}
func ExampleEqualDockerVersion() {
	// *DryRunFlag = false
	DryRunPush(MinDockerVersion)
	compareDockerVersion()
	// Output:
	// docker version --format {{.Server.Version}}
}

func ExampleInHomePath() {
	// *DryRunFlag = false
	os.Setenv("HOME", "/tmp")
	fmt.Println(inHomePath("/tmp/openwhisk-ide"))
	fmt.Println(inHomePath("/var/run"))
	// Output:
	// <nil>
	// Path /var/run is not subdir of homeDir /tmp
}
