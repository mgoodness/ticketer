// Copyright © 2018 Michael Goodness <mgoodness@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/mgoodness/ticketer/api/venues"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type venueInfo struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	City     string `json:"city,omitempty"`
	State    string `json:"state,omitempty"`
	Capacity int32  `json:"capacity,omitempty"`
}

type venueService struct {
	venues map[string]venueInfo
}

func (v *venueService) GetVenueInfo(ctx context.Context,
	r *api.GetVenueInfoRequest) (*api.VenueInfo, error) {

	// lookup venue info for ID supplied in request
	vi, ok := v.venues[r.GetId()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "venue ID %s not found", r.GetId())
	}

	// success
	return &api.VenueInfo{
		Id:       r.GetId(),
		Name:     vi.Name,
		City:     vi.City,
		State:    vi.State,
		Capacity: vi.Capacity,
	}, nil
}

func (v *venueService) ListVenues(ctx context.Context,
	r *api.ListVenuesRequest) (*api.ListVenuesResponse, error) {

	var venues []*api.VenueInfo

	// build slice with all the available venues
	for k, v := range v.venues {
		vi := &api.VenueInfo{
			Id:       k,
			Name:     v.Name,
			City:     v.City,
			State:    v.State,
			Capacity: v.Capacity,
		}

		venues = append(venues, vi)
	}

	return &api.ListVenuesResponse{
		Venues: venues,
	}, nil
}

// Run loads venue information from dataFile and starts a gRPC server
// listening on listenPort
func Run(dataFile string, listenPort int) {
	svc := &venueService{
		venues: make(map[string]venueInfo),
	}

	jsonData, err := ioutil.ReadFile(dataFile) // For read access.
	if err != nil {
		log.WithFields(log.Fields{
			"event": "startup",
			"error": err,
		}).Fatal("Failed to read data file")
	}

	//read venues from JSON data file
	err = json.Unmarshal(jsonData, &svc.venues)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "startup",
			"error": err,
		}).Fatal("Failed to marshal venue data")
	}

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(listenPort))
	if err != nil {
		log.WithFields(log.Fields{
			"event": "startup",
			"error": err,
		}).Fatal("Failed to listen")
	}

	log.WithFields(log.Fields{
		"event": "startup",
	}).Info("Listening on port ", strconv.Itoa(listenPort))
	server := grpc.NewServer()

	api.RegisterVenueServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.WithFields(log.Fields{
			"event": "startup",
			"error": err,
		}).Fatal("Failed to serve")
	}
}
