run:
	go run server.go
up:
	docker-compose -f docker-compose_env.yml up
down:
	docker-compose -f docker-compose_env.yml down
