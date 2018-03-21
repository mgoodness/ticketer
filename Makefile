.PHONY: all api build

all: build

api/api.pb.go: api/api.proto
	@protoc -I api/ \
		--go_out=plugins=grpc:api \
		api/api.proto

api: api/api.pb.go

build: api
	@go build -i -v .

clean:
	@rm api/*.pb.go ticketer
