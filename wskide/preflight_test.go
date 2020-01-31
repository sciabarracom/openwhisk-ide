package wskide

func ExampleCompareDockerVersion() {
	*DryRunFlag = false
	compareDockerVersion()
	// Output:
	// La versione di docker installata 19.3.5 è supportata
}
func ExamplePassDockerVersion() {
	// *DryRunFlag = false
	DryRunPush("19.03.5")
	compareDockerVersion()
	// Output:
	// docker version --format {{.Server.Version}}
	// La versione di docker installata 19.3.5 è supportata
}
func ExampleFailDockerVersion() {
	// *DryRunFlag = false
	DryRunPush("10.03.5")
	compareDockerVersion()
	// Output:
	// docker version --format {{.Server.Version}}
	// La versione di docker 10.3.5 installata non è piu' supportata
}
func ExampleEqualDockerVersion() {
	// *DryRunFlag = false
	DryRunPush(MinDockerVersion)
	compareDockerVersion()
	// Output:
	// docker version --format {{.Server.Version}}
	// La versione di docker installata 18.6.3-ce è supportata
}
