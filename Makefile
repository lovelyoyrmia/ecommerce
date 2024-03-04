DB_URL=postgresql://root:root@localhost:5432/ecommerce?sslmode=disable


postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_PASSWORD=root -e POSTGRES_USER=root -d postgres:12-alpine

server:
	go run cmd/main.go

seed:
	go run seeds/seed.go

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root ecommerce

dropdb:
	docker exec -it postgres12 dropdb ecommerce

migrateup:
	migrate -path migrations -database "$(DB_URL)" --verbose up

createmigrate:
	migrate create -ext sql -dir migrations -seq init_schema

migratedown:
	migrate -path migrations -database "$(DB_URL)" --verbose down

newmigration:
	migrate create -ext sql -dir migrations -seq $(name)

sqlc:
	sqlc generate

proto:
	rm -rf domain/pb/*
	rm -rf docs/swagger/*
	protoc --proto_path=domain/proto --go_out=domain/pb --go_opt=paths=source_relative \
    --go-grpc_out=domain/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=domain/pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=foedie \
    domain/proto/*.proto

dbdocs:
	dbdocs build docs/db.dbml

dbschema:
	dbml2sql --postgres -o internal/schema/schema.sql docs/db.dbml
