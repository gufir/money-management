DB_URL= postgresql://root:secret@localhost:5432/money_management?sslmode=disable

posgresql:
	docker start db_postgres_1

createdb:
	docker exec -it  db_postgres_1 createdb --username=root --owner=root money_management

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

migration:
	migrate create -ext sql -dir db/migration -seq init_schema

.PHONY: posgresql sqlc createdb migrateup migratedown migration