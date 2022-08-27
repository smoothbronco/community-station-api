test:
	go test -v ./...

gRPC-dev:
	go run location/server/server.go

GraphQL-dev:
	go run graph/server/server.go
