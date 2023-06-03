stop-postgres-podman:
	podman-compose -f docker-compose.yaml down

delete-postgres-podman:
	podman-compose -f docker-compose.yaml down -v

start-postgres-podman:
	podman-compose -f docker-compose.yaml up -d

create-db-podman:
	podman exec -it postgres_container createdb --username=postgres --owner=postgres bookshelf

drop-db-podman:
	podman exec -it postgres_container dropdb --username=postgres bookshelf

migrate-up:
	migrate -path db/migration -database "postgresql://postgres:postgres@127.0.0.1:5432/bookshelf?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://postgres:postgres@127.0.0.1:5432/bookshelf?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: stop-postgres-podman delete-postgres-podman start-postgres-podman create-db-podman drop-db-podman migrate-up migrate-down sqlc