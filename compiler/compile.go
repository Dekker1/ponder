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

// CompileDir compiles all lilypond files and makes all
// sheet music available in the OutputDir
func CompileDir(path string, opts *settings.Settings) {

}

func generateScores() ([]*settings.Score, func(string, os.FileInfo) error) {
	scores := make([]*settings.Score)
	return scores, func(path string, file os.FileInfo) error {
		switch filepath.Ext(path) {
		case ".ly":
			log.WithFields(log.Fields{"path": path}).Info("adding lilypond file")
			append(scores, &settings.Score{Path: path})

		case ".json":
			if filepath.Base(path) != "ponder.json" {
				log.WithFields(log.Fields{"path": path}).Info("adding json file")
				if score, err := fromJSON(path); err != nil {
					log.WithFields(log.Fields{
						"error": err,
						"path":  path,
					}).Warning("unable to parse score settings, skipping...")
				} else {
					append(scores, score)
				}
			}

		default:
			log.WithFields(log.Fields{"path": path}).Debug("ignoring file")
		}
	}
}
