up: docker_up
down: docker_down
pull: docker_pull
build: docker_build
ps: docker_ps
restart: down up

run: go_run

m: migrate

migrate:
	./migrate.sh

test:
	go test -v ./...

docker_up:
	 docker compose up -d

docker_down:
	docker compose down --remove-orphans

docker_pull:
	 docker compose pull

docker_build:
	 docker compose build --pull

docker_ps:
	docker compose ps

# GO

go_build:
	go build ./cmd/kpi

go_run:
	./run.sh
