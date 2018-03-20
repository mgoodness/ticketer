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

type venueInfo struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

/* venueService implements VenueServiceServer as defined in the generated code:
// Server API for VenueService service

type VenueServiceServer interface {
	GetVenueInfo(context.Context, *GetVenueInfoRequest) (*VenueInfo, error)
	ListVenues(context.Context, *ListVenuesRequest) (*ListVenuesResponse, error)
}
*/
type venueService struct {
	venues map[string]venueInfo
}

func (v *venueService) GetVenueInfo(ctx context.Context,
	r *api.GetVenueInfoRequest) (*api.VenueInfo, error) {

	// lookup venue info for ID supplied in request
	vi, ok := v.venues[r.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "Venue ID %s not found", r.GetId())
	}

	// success
	return &api.VenueInfo{
		Id:    r.GetId(),
		Name:  vi.Name,
		City:  vi.City,
		State: vi.State,
	}, nil
}

func (v *venueService) ListVenues(ctx context.Context,
	r *api.ListVenuesRequest) (*api.ListVenuesResponse, error) {

	var venues []*api.VenueInfo

	// build slice with all the available venues
	for k, v := range v.venues {
		vi := &api.VenueInfo{
			Id:    k,
			Name:  v.Name,
			City:  v.City,
			State: v.State,
		}

		venues = append(venues, vi)
	}

	return &api.ListVenuesResponse{
		Venues: venues,
	}, nil
}

// Run the server
func Run(listenPort int, dataFile string) {
	svc := &venueService{
		venues: make(map[string]venueInfo),
	}

	jsonData, err := ioutil.ReadFile(dataFile) // For read access.
	if err != nil {
		log.Fatalf("Failed to read data file: %v", err)
	}

	//read venues from JSON data file
	err = json.Unmarshal(jsonData, &svc.venues)
	if err != nil {
		log.Fatalf("Failed to marshal venue data: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(listenPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("Listening on", strconv.Itoa(listenPort))
	server := grpc.NewServer()

	api.RegisterVenueServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
