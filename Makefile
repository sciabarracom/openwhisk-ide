BASE=actionloop

.PHONY: help
help: 
	@echo TODO

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
	actionloop/standalone

#	-e "CONFIG_FORCE_WHISK_CONTAINER__FACTORY_CONTAINER__ARGS_EXTRA__ARGS_ENV=[__OW_DEBUG_PORT=8081]" \
#	-e 'JVM_EXTRA_ARGS=-Dwhisk.container-factory.container-args.extra-args.env=["__OW_DEBUG_PORT=8081"]' \
.PHONY: ide
ide:
	docker build -t actionlooop/ide-js ide-js
	docker run -ti \
	--rm --name ide-js \
	-p 3000:3000 \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-v $(PWD)/project:/home/project \
	--add-host openwhisk:$(shell docker inspect openwhisk | jq -r '.[0].NetworkSettings.IPAddress') \
	actionloop/ide-typescript

