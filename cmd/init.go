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
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/jjdekker/ponder/helpers"
	"github.com/spf13/cobra"
)

var (
	settingsFile     = "ponder.json"
	settingsTemplate = []byte(`{
  "Name": "",
  "IgnoreDirs": [".git"],
  "LilypondIncludes": [],
  "OutputDir": "out"
}`)
	gitIgnoreTemplate = []byte(`# Output Folder
out/`)
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [name]",
	Short: "Initialize a Ponder Library",
	Long: `Initialize (ponder init) will create a new library, with a ponder
	settings file and corresponding git ignore file.

  * If a name is provided, it will be created in the current directory;
  * If no name is provided, the current directory will be assumed;
Init will not use an existing directory with contents.`,
	Run: func(cmd *cobra.Command, args []string) {
		var path string
		var err error
		switch len(args) {
		case 0:
			path, err = helpers.CleanPath("")
		case 1:
			path, err = helpers.CleanPath(args[0])
		default:
			log.Fatal("init command does not support more than 1 parameter")
		}
		helpers.Check(err, "Given path is invalid")

		initializePath(path)
	},
}

func initializePath(path string) {
	b, err := helpers.Exists(path)
	if err != nil {
		log.WithFields(log.Fields{"error": err, "path": path}).
			Fatal("Unable to check path")
	}

	if !b {
		err := os.MkdirAll(path, os.ModePerm)
		helpers.Check(err, "Could not create directory")
	}

	createFile(filepath.Join(path, settingsFile), settingsTemplate)
	createFile(filepath.Join(path, ".gitignore"), gitIgnoreTemplate)
}

func createFile(path string, content []byte) {
	b, err := helpers.Exists(path)
	if err != nil {
		log.WithFields(log.Fields{"error": err, "path": path}).
			Fatal("Unable to check path")
	}

	if !b {
		err = ioutil.WriteFile(path, content, 0644)
		if err != nil {
			log.WithFields(log.Fields{"error": err, "path": path}).
				Fatal("Unable to create settings file")
		}
	}
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
