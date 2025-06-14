.PHONY: build_hash

build_hash:
	@protoc -I proto proto/hash.proto --go_out=./pkg/grpc --go-grpc_out=./pkg/grpc