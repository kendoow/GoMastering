postgres:
	docker run --name postgresbank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
postgresstart:
	docker start postgresbank
createdb:
	docker exec -it postgresbank createdb --username=root --owner=root simple_bank
dropdb:	
	docker exec -it postgresbank dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up	
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down	
test:
	go test -v -cover ./...	
.PHONY:postgres postgresstart createdb dropdb migrateup migratedown test