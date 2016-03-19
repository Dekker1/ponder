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

package compiler

import (
	"os"
	"os/exec"

	"github.com/jjdekker/ponder/helpers"
	"github.com/jjdekker/ponder/settings"
)

var (
	lilypondCmd  = "lilypond"
	lilypondArgs []string
)

// PrepareLilypond sets all arguments and options for the Lilypond
// compilation function using the given settings
func PrepareLilypond(opts *settings.Settings) {
	// Adds all includes to the lilypond arguments
	for _, dir := range opts.LilypondIncludes {
		lilypondArgs = append(lilypondArgs, "--include=\""+dir+"\"")
	}
	lilypondArgs = append(lilypondArgs, "--loglevel=ERROR")
	lilypondArgs = append(lilypondArgs, "--pdf")

	// TODO: Make this an absolute path.
	lilypondArgs = append(lilypondArgs, "--output=\""+opts.OutputDir+"\"")
	if !helpers.Exists(opts.OutputDir) {
		err := os.MkdirAll(opts.OutputDir, os.ModePerm)
		helpers.Check(err, "Could not create output directory")
	}
}

// Lilypond runs the lilypond compiler on the given path
// using the arguments prepared by the PrepareLilypond function
func Lilypond(path string) (string, error) {
	cmd := exec.Command(lilypondCmd, append(lilypondArgs, path)...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
