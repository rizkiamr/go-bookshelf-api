stop-postgres-podman:
	podman-compose -f docker-compose.yaml down postgres

delete-postgres-podman:
	podman-compose -f docker-compose.yaml down postgres -v

start-postgres-podman:
	podman-compose -f docker-compose.yaml up postgres -d

stop-pgadmin-podman:
	podman-compose -f docker-compose.yaml down pgadmin

delete-pgadmin-podman:
	podman-compose -f docker-compose.yaml down pgadmin -v

start-pgadmin-podman:
	podman-compose -f docker-compose.yaml up -d pgadmin

create-db-podman:
	podman exec -it postgres_container createdb --username=postgres --owner=postgres bookshelf

drop-db-podman:
	podman exec -it postgres_container dropdb --username=postgres bookshelf

stop-postgres-docker:
	docker compose -f docker-compose.yaml down postgres

delete-postgres-docker:
	docker compose -f docker-compose.yaml down postgres -v

start-postgres-docker:
	docker compose -f docker-compose.yaml up postgres -d

stop-pgadmin-docker:
	docker compose -f docker-compose.yaml down pgadmin

delete-pgadmin-docker:
	docker compose -f docker-compose.yaml down pgadmin -v

start-pgadmin-docker:
	docker compose -f docker-compose.yaml up -d pgadmin

create-db-docker:
	docker exec -it postgres_container createdb --username=postgres --owner=postgres bookshelf

drop-db-docker:
	docker exec -it postgres_container dropdb --username=postgres bookshelf

migrate-up:
	migrate -path db/migration -database "postgresql://postgres:postgres@127.0.0.1:5432/bookshelf?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://postgres:postgres@127.0.0.1:5432/bookshelf?sslmode=disable" -verbose down

sqlc:
	sqlc generate

clean-test-cache:
	go clean -testcache

test:
	go clean -testcache; go test -v -cover ./...

server:
	go run main.go

.PHONY: stop-postgres-docker delete-postgres-docker start-postgres-docker stop-pgadmin-docker delete-pgadmin-docker start-pgadmin-docker create-db-docker drop-db-docker stop-postgres-podman delete-postgres-podman start-postgres-podman stop-pgadmin-podman delete-pgadmin-podman start-pgadmin-podman create-db-podman drop-db-podman migrate-up migrate-down sqlc clean-test-cache test server