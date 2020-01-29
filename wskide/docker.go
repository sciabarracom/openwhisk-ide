package wskide

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

func dockerRunIde() string {
	//--add-host openwhisk:$(docker inspect openwhisk | jq -r '.[0].NetworkSettings.IPAddress')
	return Sys(`docker run -d -p 3000:3000
--rm --name ide-js
-v /var/run/docker.sock:/var/run/docker.sock
-v /${PWD}/project:/home/project
actionloop/ide-js`)
}

func dockerRmIde() string {
	return Sys("docker exec ide-js stop")
}
