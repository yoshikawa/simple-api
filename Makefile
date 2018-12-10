DB_CONTAINER_NAME:=simple-mysql
DBNAME:=api

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/up:
	docker-compose exec api goose up

migrate/down:
	docker-compose exec api goose down

docker/build:
	docker-compose build

docker/start:
	docker-compose up -d

docker/logs:
	docker-compose logs

docker/stop:
	docker-compose stop

docker/clean:
	docker-compose rm

api/bash:
	docker-compose exec api bash

db/bash:
	docker-compose exec db bash

api/init:
	docker-compose exec api dep ensure

run:
	docker-compose exec api go run main.go

doc:
	godoc -http=:6060