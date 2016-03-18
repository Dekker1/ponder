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

import "github.com/jjdekker/ponder/settings"

var (
	lilypondCmd  = "lilypond"
	lilypondArgs []string
)

//
func prepareLilypondArgs(opts *settings.Settings) {
	// Adds all includes to the lilypond arguments
	for dir := range opts.LilypondIncludes {
		lilypondArgs = append(lilypondArgs, "--include=\""+dir+"\"")
	}
	lilypondArgs = append(lilypondArgs, "--loglevel=ERROR")
	lilypondArgs = append(lilypondArgs, "--pdf")
	// TODO: Make this an absolute path.
	lilypondArgs = append(lilypondArgs, "--output=\""+opts.OutputDir+"\"")
}

// Lilypond runs the lilypond compiler
func Lilypond(path string) error {
	return nil
}
