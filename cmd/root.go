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
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	veryVerbose bool
	verbose     bool
)

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ponder",
	Short: "A managing tool for lilypond sheet music libraries",
	Long: `Ponder is a tool to help manage your sheet music library.
The main purpose is to help in the compilation of your lilypond files
into both single files and a fully functioning song book. It also accepts
other PDF files to be part of your song book.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output of application events")
	RootCmd.PersistentFlags().BoolVar(&veryVerbose, "vv", false, "Debug output of application events")

	setLogLevel()
}

func setLogLevel() {
	if veryVerbose {
		log.SetLevel(log.DebugLevel)
	} else if verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}
