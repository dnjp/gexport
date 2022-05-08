.PHONY: proto
proto:
	cd proto && \
		protoc \
		--go_out=gauth --go_opt=paths=source_relative \
		--go-grpc_out=gauth --go-grpc_opt=paths=source_relative \
		gauth.proto

install: proto
	go install main.go

build: proto
	mkdir -p bin
	go build -o bin/gexport ./main.go
