include .env
local:
	go run server.go -f config_local.yml
env:
	docker run --rm --name $(MONGO_CONTAINER_NAME) -p $(MONGO_HOST_PORT):27017 -v /data/db:/data/db -e MONGO_INITDB_ROOT_USERNAME=$(MONGO_ROOT) -e MONGO_INITDB_ROOT_PASSWORD=$(MONGO_ROOT_PW) -d mongo:4.4

up:
	docker-compose up
down:
	docker-compose down
