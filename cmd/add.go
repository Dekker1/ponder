// Copyright © 2016 Jip J. Dekker <jip@dekker.li>
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
	"log"

	"github.com/jjdekker/ponder/helpers"
	"github.com/jjdekker/ponder/settings"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [file]",
	Short: "Add pdf file to a book",
	Long: `Add creates a json file with all options regarding a sheet music file in PDF format.
The information saved in the json file will be used when compiling the songbook.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			path string
			err  error
		)
		switch len(args) {
		case 1:
			path, err = helpers.CleanPath(args[0])
			helpers.Check(err, "Unable to create valid path")
		default:
			log.Fatal("the add command needs exactly 1 parameter")
		}
		dir, _ := getSettings()
		settings.CreateScore(path, dir)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
