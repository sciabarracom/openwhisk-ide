BASE=actionloop

.PHONY: help
help: 
	@echo first, make openwhisk then make ide

.PHONY: clean
clean:
	-docker ps -qa | xargs docker kill
	-docker ps -qa | xargs docker rm

.PHONY: openwhisk
openwhisk:
	docker run -ti \
	-p 3232:3232 -p 3233:3233 \
	--rm --name openwhisk --hostname openwhisk \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-e CONTAINER_EXTRA_ENV=__OW_DEBUG_PORT=8081 \
	openwhisk/standalone:nightly

.PHONY: ide-base
ide-base:
	docker build -t actionloop/ide-js-base ide-js-base

.PHONY: ide
ide:
	docker build -t actionloop/ide-js ide-js-poc
	docker run -ti \
	--rm --name ide-js \
	-p 3000:3000 \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-v $(PWD)/project:/home/project \
	--add-host openwhisk:$(shell docker inspect openwhisk | jq -r '.[0].NetworkSettings.IPAddress') \
	actionloop/ide-js

