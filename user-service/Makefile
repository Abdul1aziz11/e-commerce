export POSTGRES_DB=postgres://postgres:1234@localhost:5432/gatewaydb?sslmode=disable

gen-user:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./protos/user-service/user.proto

umigfile:
	migrate create -ext sql -dir migrations/ user-service

uup:
	migrate -path migrations -database $(POSTGRES_DB) up

udown:
	migrate -path migrations -database $(POSTGRES_DB) down

uforce:
	migrate -path migrations -database $(POSTGRES_DB) force $(version)	

