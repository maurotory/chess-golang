build:
	go build -o bin/chess cmd/client/main.go
proto:
	protoc -I . api/main.proto --go-grpc_out=. --go_out=.	
run-server:
	go run cmd/server/main.go
run-client:
	go run cmd/client/main.go