package client

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/mgoodness/ticketer/api"
	"google.golang.org/grpc"
)

type byID []*api.VenueInfo

func (b byID) Len() int           { return len(b) }
func (b byID) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byID) Less(i, j int) bool { return b[i].GetId() < b[j].GetId() }

// List resources
func List(targetHost string, targetPort int) {
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
		sort.Sort(byID(venues))

		fmt.Printf("ID\tName\tCity\tState\n")
	} else {
		fmt.Println("No venues found")
	}

	for _, vi := range venues {
		fmt.Printf("%s\t%s\t%s\t%s\n",
			vi.GetId(),
			vi.GetName(),
			vi.GetCity(),
			vi.GetState())
	}
}
