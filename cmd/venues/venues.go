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

package main

import (
	"github.com/mgoodness/ticketer/internal/app/venues"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var (
	dataFile   string
	listenPort int
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	flag.StringVarP(&dataFile, "filename", "f", "", "path to data file (required)")
	flag.IntVarP(&listenPort, "listen", "l", 8080, "TCP port on which to listen")
}

func main() {
	flag.Parse()

	if len(dataFile) == 0 {
		log.WithField("event", "startup").Fatal("Data file not specified")
	}

	server.Run(dataFile, listenPort)
}
