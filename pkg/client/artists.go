package client

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/mgoodness/ticketer/api/artists"
	"google.golang.org/grpc"
)

type byArtistID []*api.ArtistInfo

func (b byArtistID) Len() int           { return len(b) }
func (b byArtistID) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byArtistID) Less(i, j int) bool { return b[i].GetId() < b[j].GetId() }

// GetArtistInfo retrieves information about a single artist from a gRPC server
func GetArtistInfo(targetHost string, targetPort int, id string) {
	conn, err := grpc.Dial(targetHost+":"+strconv.Itoa(targetPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := api.NewArtistServiceClient(conn)
	ai, err := client.GetArtistInfo(context.Background(), &api.GetArtistInfoRequest{Id: id})
	if err != nil {
		log.Fatalf("Could not get: %v", err)
	}

	fmt.Printf("ID\tName\n")
	fmt.Printf("%s\t%s\n",
		ai.GetId(),
		ai.GetName())
}

// ListArtists retrieves a list of artists from a gRPC server
func ListArtists(targetHost string, targetPort int) {
	conn, err := grpc.Dial(targetHost+":"+strconv.Itoa(targetPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := api.NewArtistServiceClient(conn)
	rsp, err := client.ListArtists(context.Background(), &api.ListArtistsRequest{})
	if err != nil {
		log.Fatalf("Could not list: %v", err)
	}

	artists := rsp.GetArtists()
	if len(artists) > 0 {
		sort.Sort(byArtistID(artists))

		fmt.Printf("ID\tName\n")
	} else {
		fmt.Println("No artists found")
	}

	for _, ai := range artists {
		fmt.Printf("%s\t%s\n",
			ai.GetId(),
			ai.GetName())
	}
}
