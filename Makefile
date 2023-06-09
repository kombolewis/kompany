.PHONY: migrateup migratedown test mock

migrateup:
	@migrate -path db/migration -database "postgres://root:secret@localhost:5432/kompani?sslmode=disable" -verbose up

migratedown:
	@migrate -path db/migration -database "postgres://root:secret@localhost:5432/kompani?sslmode=disable" -verbose down
test:
	@go test -v --cover ./...
mock:
	@mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/kombolewis/kompani/db/sqlc Store