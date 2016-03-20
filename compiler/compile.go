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
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/jjdekker/ponder/settings"
)

var (
	scores []*settings.Score
)

// CompileDir compiles all lilypond files and makes all
// sheet music available in the OutputDir
func CompileDir(path string, opts *settings.Settings) {
	// Find all scores
	collector := generateScores()
	filepath.Walk(path, compilePath(path, opts, collector))

	PrepareLilypond(opts)
	for _, score := range scores {
		// TODO: compile only if there are changes
		msg, err := Lilypond(score.Path)
		if err != nil {
			log.WithFields(log.Fields{
				"message": msg,
				"error":   err,
				"score":   score,
			}).Warning("score failed to compile")
		}
	}
}

func generateScores() func(string, os.FileInfo) error {
	return func(path string, file os.FileInfo) error {
		switch filepath.Ext(path) {
		case ".ly":
			log.WithFields(log.Fields{"path": path}).Info("adding lilypond file")
			scores = append(scores, &settings.Score{
				Path:         path,
				LastModified: file.ModTime(),
			})

		case ".json":
			if filepath.Base(path) != "ponder.json" {
				log.WithFields(log.Fields{"path": path}).Info("adding json file")
				if score, err := settings.FromJSON(path); err != nil {
					log.WithFields(log.Fields{
						"error": err,
						"path":  path,
					}).Warning("unable to parse score settings, skipping...")
				} else {
					scores = append(scores, score)
				}
			}

		default:
			log.WithFields(log.Fields{"path": path}).Debug("ignoring file")
		}
		return nil
	}
}
