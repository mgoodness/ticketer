.PHONY: all api build

all: build

api/venues.pb.go: api/venues.proto
	@protoc -I api/ \
		-I${GOPATH}/src \
		--go_out=plugins=grpc:api \
		api/venues.proto

api: api/venues.pb.go

dep:
	@dep ensure

build: dep api
	@go build -i -v .

clean:
	@rm api/*.pb.go
