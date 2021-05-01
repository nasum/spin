db-up:
	migrate -database $(DATABASE_URL) -path migrations up
db-down:
	migrate -database $()DATABASE_URL) -path migrations down
test:
	go test -v ./usecase ./entity
