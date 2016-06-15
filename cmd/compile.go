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
	"github.com/jjdekker/ponder/compiler"
	"github.com/spf13/cobra"
)

// compileCmd represents the compile command
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compiles all lilypond files in the library",
	Long: `Compile (ponder compile) will walk through and compile all
lilypond files in accordance to ponder settings file.
Files that have already been compiled will be skipped,
unless the lilypond file has been edited.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, opts := getSettings()
		opts.EnablePointAndClick = opts.EnablePointAndClick || enablePointAndClick
		compiler.CompileDir(path, opts)
	},
}

func init() {
	compileCmd.Flags().BoolVarP(&enablePointAndClick, "clickable",
		"c", false, "Enable Lilypond's Point and Click functionality")
	RootCmd.AddCommand(compileCmd)
}
