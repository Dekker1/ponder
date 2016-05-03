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
	"os/exec"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
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
		lilypondArgs = append(lilypondArgs, "--include="+dir)
	}
	lilypondArgs = append(lilypondArgs, "--loglevel=ERROR")
	lilypondArgs = append(lilypondArgs, "--pdf")
}

// Lilypond runs the lilypond compiler on the given path
// using the arguments prepared by the PrepareLilypond function
func Lilypond(s *settings.Score) (string, error) {
	args := append(lilypondArgs, "--output="+filepath.Dir(s.OutputPath))
	err := helpers.ExistsOrCreate(filepath.Dir(s.OutputPath))
	if err != nil {
		return "", err
	}
	args = append(args, s.Path)

	cmd := exec.Command(lilypondCmd, args...)
	log.WithFields(log.Fields{
		"path": s.Path,
		"cmd":  cmd,
	}).Info("compiling file using lilypond")
	out, err := cmd.CombinedOutput()
	log.Debug("finished compiling")
	return string(out), err
}
