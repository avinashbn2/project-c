run:
	go run cmd/app/*.go
postgres:
	docker run --rm -ti --network host -e POSTGRES_PASSWORD=secret postgres
adminer:
	docker run --rm -ti --network host adminer
migrate:
	migrate -source file://migrations -verbose -database postgres://postgres:changeme@localhost:5432/postgres?sslmode=disable up

migratedown:
	migrate -source file://migrations -verbose -database postgres://postgres:changeme@localhost:5432/postgres?sslmode=disable down
