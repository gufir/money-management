DB_URL= postgresql://root:secret@localhost:5432/money_management?sslmode=disable

posgresql:
	docker start db_postgres_1

createdb:
	docker exec -it  db_postgres_1 createdb --username=root --owner=root money_management

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

dropdb:
	docker exec -it  db_postgres_1 dropdb --username=root money_management

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

migration:
	migrate create -ext sql -dir db/migration -seq init_schema

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/gufir/money-management/db/sqlc Store

run:
	go run main.go

proto:
	rm -rf pb/*.go
	protoc --proto_path=proto \
		--go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--experimental_allow_proto3_optional \
		proto/*.proto

.PHONY: server posgresql sqlc createdb migrateup migratedown migration dropdb mock new_migration proto