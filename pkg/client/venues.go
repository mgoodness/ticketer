package client

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/mgoodness/ticketer/api/venues"
	"google.golang.org/grpc"
)

type byVenueID []*api.VenueInfo

func (b byVenueID) Len() int           { return len(b) }
func (b byVenueID) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byVenueID) Less(i, j int) bool { return b[i].GetId() < b[j].GetId() }

// GetVenueInfo retrieves information about a single venue from a gRPC server
func GetVenueInfo(targetHost string, targetPort int, id string) {
	conn, err := grpc.Dial(targetHost+":"+strconv.Itoa(targetPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := api.NewVenueServiceClient(conn)
	vi, err := client.GetVenueInfo(context.Background(), &api.GetVenueInfoRequest{Id: id})
	if err != nil {
		log.Fatalf("Could not get: %v", err)
	}

	fmt.Printf("ID\tName\tCity\tState\tCapacity\n")
	fmt.Printf("%s\t%s\t%s\t%s\t%d\n",
		vi.GetId(),
		vi.GetName(),
		vi.GetCity(),
		vi.GetState(),
		vi.GetCapacity())
}

// ListVenues retrieves a list of venues from a gRPC server
func ListVenues(targetHost string, targetPort int) {
	conn, err := grpc.Dial(targetHost+":"+strconv.Itoa(targetPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := api.NewVenueServiceClient(conn)
	rsp, err := client.ListVenues(context.Background(), &api.ListVenuesRequest{})
	if err != nil {
		log.Fatalf("Could not list: %v", err)
	}

	venues := rsp.GetVenues()
	if len(venues) > 0 {
		sort.Sort(byVenueID(venues))

		fmt.Printf("ID\tName\tCity\tState\tCapacity\n")
	} else {
		fmt.Println("No venues found")
	}

	for _, vi := range venues {
		fmt.Printf("%s\t%s\t%s\t%s\t%d\n",
			vi.GetId(),
			vi.GetName(),
			vi.GetCity(),
			vi.GetState(),
			vi.GetCapacity())
	}
}
