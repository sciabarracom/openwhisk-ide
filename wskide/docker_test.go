package wskide

import "fmt"

func ExampleDockerVersion() {
	// *DryRunFlag = false
	DryRunPush("19.03.5")
	fmt.Println(dockerVersion())
	// Output:
	// docker version --format {{.Server.Version}}
	// 19.03.5
}
func ExampleDockerRunOpenWhisk() {
	// *DryRunFlag = false
	DryRunPush("991c7972fa4612c873b3804a4c334b3af66687a7f1e548a36dfdfe0c6a717cbe")
	fmt.Println(dockerRunOpenWhisk())
	// Output:
	// docker run -d -p 3232:3232 -p 3233:3233 --rm --name openwhisk --hostname openwhisk -v /var/run/docker.sock:/var/run/docker.sock -e CONTAINER_EXTRA_ENV=__OW_DEBUG_PORT=8081 openwhisk/standalone:nightly
	// 991c7972fa4612c873b3804a4c334b3af66687a7f1e548a36dfdfe0c6a717cbe
}
func ExampleDockerRmOpenWhisk() {
	// *DryRunFlag = false
	DryRunPush()
	fmt.Println(dockerRmOpenWhisk())
	//  Output:
	//  docker exec openwhisk stop
}
func ExampleDockerRunIde() {
	// *DryRunFlag = false
	DryRunPush("641792b3e0112c8fa1896b8944a846dbbab88fe5729f3d464e71475afd9e6057")
	fmt.Println(dockerRunIde())
	// Output:
	// docker run -d -p 3000:3000 --rm --name ide-js -v /var/run/docker.sock:/var/run/docker.sock -v /${PWD}/project:/home/project actionloop/ide-js
	// 641792b3e0112c8fa1896b8944a846dbbab88fe5729f3d464e71475afd9e6057
}
func ExampleDockerRmIde() {
	// *DryRunFlag = false
	DryRunPush()
	fmt.Println(dockerRmIde())
	// Output:
	// docker exec ide-js stop
}
