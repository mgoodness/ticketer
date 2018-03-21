.PHONY: all api build_artists build_tictl build_venues dep

all: build_artists build_tictl build_venues

api/artists/api.pb.go: api/artists/api.proto
	@protoc -I api/artists/ \
		--go_out=plugins=grpc:api/artists \
		api/artists/api.proto

api/venues/api.pb.go: api/venues/api.proto
	@protoc -I api/venues/ \
		--go_out=plugins=grpc:api/venues \
		api/venues/api.proto

api: api/venues/api.pb.go api/artists/api.pb.go

build_artists: api
	@go build -i -v -o bin/artists ./cmd/artists

build_tictl: api
	@go build -i -v -o bin/tictl ./cmd/tictl

build_venues: api
	@go build -i -v -o bin/venues ./cmd/venues

clean:
	@rm api/{artists,venues}/*.pb.go bin/*

dep:
	@dep ensure --update
