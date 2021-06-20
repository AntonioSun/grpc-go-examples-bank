.PHONY: compile
compile: ## Compile the proto file.
	protoc -I pkg/proto/bank/account/ pkg/proto/bank/account/deposit.proto --go_out=plugins=grpc:pkg/proto/bank/account/

.PHONY: run
run: ## Build and run server.
	go build -race -ldflags "-s -w" -o bin/server cmd/server/main.go
	bin/server