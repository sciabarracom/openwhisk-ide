package wskide

import (
	"fmt"
	"path/filepath"
	"strings"
)

func dockerVersion() string {
	return Sys("@docker version --format {{.Server.Version}}")
}

func dockerImages() string {
	return Sys("docker images")
}

func dockerRunOpenWhisk() string {
	return Sys(`docker run -d -p 3232:3232
-p 3233:3233 --rm --name openwhisk
--hostname openwhisk -v /var/run/docker.sock:/var/run/docker.sock
-e CONTAINER_EXTRA_ENV=__OW_DEBUG_PORT=8081 openwhisk/standalone:nightly`)
}

func dockerRmOpenWhisk() string {
	return Sys("docker exec openwhisk stop")
}

func dockerRunIde(dir string) error {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}
	openwhiskIP := Sys("docker inspect", "--format={{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}", "openwhisk")
	if strings.HasPrefix(openwhiskIP, "Error:") {
		return fmt.Errorf("%s", openwhiskIP)
	}
	command := fmt.Sprintf(`docker run -d -p 3000:3000
--rm --name ide-js
-v /var/run/docker.sock:/var/run/docker.sock
-v %s:/home/project
actionloop/ide-js --add-host %s`, dir, openwhiskIP)
	Sys(command)
	return nil
}

func dockerRmIde() string {
	return Sys("docker kill ide-js")
}
