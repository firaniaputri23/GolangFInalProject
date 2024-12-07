postgresinit:
	docker run --name postgres -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

postgres:
	docker exec -it postgres psql

createdb:
	docker exec -it postgres createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it postgres dropdb go-chat

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown
