# Go Lang

## Install Dependencies
- go get -u github.com/golang/protobuf/{proto,protoc-gen-go} 

## Generate Proto Buffers
- protoc --proto_path=proto/ proto/*.proto --plugin=/Users/jamesson/go/bin/protoc-gen-go-rpc  --go_out=. --go-grpc_out=.

## Run client
- go run cmd/client/client.go

## Run server
-  go run cmd/server/server.go