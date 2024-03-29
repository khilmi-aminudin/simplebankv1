DB_SOURCE=postgresql://root:secret@localhost:5432/simplebank?sslmode=disable

postgresql :
	docker run --name postgres-simplebank --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

execdb :
	docker exec -it postgres-simplebank psql simplebank

createdb :
	docker exec -it postgres-simplebank createdb --username=root --owner=root simplebank

rundb :
	docker start postgres-simplebank

initmigrate :
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup :
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up

migratedown :
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down

migrateup1 :
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose up 1

migratedown1 :
	migrate -path db/migration -database "$(DB_SOURCE)" -verbose down 1

sqlc :
	sqlc generate

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

test :
	go test -v -cover ./...

runserver :
	go run main.go

mock :
	mockgen -package mockdb -destination db/mock/store.go github.com/khilmi-aminudin/simplebankv1/db/sqlc Store

.PHONY : postgresql execdb createdb initmigrate migrateup migratedown migrateup1 migratedown1 sqlc db_docs db_schema test runserver mock