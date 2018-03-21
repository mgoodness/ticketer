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

	"github.com/mgoodness/ticketer/api/artists"
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

// Run loads artist information from dataFile and starts a gRPC server
// listening on listenPort
func Run(dataFile string, listenPort int) {
	svc := &artistService{
		artists: make(map[string]artistInfo),
	}

	jsonData, err := ioutil.ReadFile(dataFile) // For read access.
	if err != nil {
		log.WithFields(log.Fields{
			"event": "startup",
			"error": err,
		}).Fatal("Failed to read data file")
	}

	//read artists from JSON data file
	err = json.Unmarshal(jsonData, &svc.artists)
	if err != nil {
		log.WithFields(log.Fields{
			"event": "startup",
			"error": err,
		}).Fatal("Failed to marshal artist data")
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

	api.RegisterArtistServiceServer(server, svc)
	if err := server.Serve(lis); err != nil {
		log.WithFields(log.Fields{
			"event": "startup",
			"error": err,
		}).Fatal("Failed to serve")
	}
}
