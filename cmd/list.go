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

package cmd

import (
	"errors"

	"github.com/mgoodness/ticketer/pkg/venues/client"
	"github.com/spf13/cobra"
)

var (
	targetPort int
	targetHost string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List resources over gRPC",
	Long:  "Lists resources over gRPC.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 || len(args) > 1 || args[0] != "venues" {
			return errors.New("exactly one valid resource type is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.List(targetHost, targetPort)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	listCmd.PersistentFlags().StringVar(&targetHost, "host", "localhost", "target host")
	listCmd.PersistentFlags().IntVar(&targetPort, "port", 8080, "target port")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
