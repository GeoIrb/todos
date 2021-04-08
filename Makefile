PWD                = $(shell pwd)

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    pkg/user/rpc/rpc.proto

lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint golangci-lint run -v -E maligned 
