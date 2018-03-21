package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"strconv"

	"github.com/mgoodness/ticketer/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type artistInfo struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type artistService struct {
	artists map[string]artistInfo
}

func (v *artistService) GetArtistInfo(ctx context.Context,
	r *api.GetArtistInfoRequest) (*api.ArtistInfo, error) {

	// lookup artist info for ID supplied in request
	vi, ok := v.artists[r.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "artist ID %s not found", r.GetId())
	}

	// success
	return &api.ArtistInfo{
		Id:   r.GetId(),
		Name: vi.Name,
	}, nil
}

func (v *artistService) ListArtists(ctx context.Context,
	r *api.ListArtistsRequest) (*api.ListArtistsResponse, error) {

	var artists []*api.ArtistInfo

	// build slice with all the available artists
	for k, v := range v.artists {
		vi := &api.ArtistInfo{
			Id:   k,
			Name: v.Name,
		}

		artists = append(artists, vi)
	}

	return &api.ListArtistsResponse{
		Artists: artists,
	}, nil
}

// RunArtistServer loads artist information from dataFile and starts a gRPC server
// listening on listenPort
func RunArtistServer(listenPort int, dataFile string) {
	svc := &artistService{
		artists: make(map[string]artistInfo),
	}

	jsonData, err := ioutil.ReadFile(dataFile) // For read access.
	if err != nil {
		log.Fatalf("Failed to read data file: %v", err)
	}

	//read artists from JSON data file
	err = json.Unmarshal(jsonData, &svc.artists)
	if err != nil {
		log.Fatalf("Failed to marshal artist data: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(listenPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Listening on", strconv.Itoa(listenPort))
	server := grpc.NewServer()

	api.RegisterArtistServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
