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

var keepTemplate bool
var enablePointAndClick bool

// bookCmd represents the book command
var bookCmd = &cobra.Command{
	Use:   "book",
	Short: "Generate target songbook",
	Long: `Generates a songbook based on a target specified in the ponder
settings. If not all songs included in the target have been compiled,
these songs will be compiled as if the compile command had been called.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, opts := getSettings()
		opts.KeepBookTemplate = opts.KeepBookTemplate || keepTemplate
		opts.EnablePointAndClick = opts.EnablePointAndClick || enablePointAndClick
		compiler.MakeBook(path, opts)
	},
}

func init() {
	bookCmd.Flags().BoolVarP(&keepTemplate, "keep-template",
		"k", false, "Leave the LaTeX source for the book in the output directory")
	bookCmd.Flags().BoolVarP(&enablePointAndClick, "clickable",
		"c", false, "Enable Lilypond's Point and Click functionality")
	RootCmd.AddCommand(bookCmd)
}
