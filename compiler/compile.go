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
	"github.com/jjdekker/ponder/helpers"
	"github.com/jjdekker/ponder/settings"
)

var (
	scores []settings.Score
)

// CompileDir compiles all lilypond files and makes all
// sheet music available in the OutputDir
func CompileDir(path string, opts *settings.Settings) {
	// Find all scores
	collector := generateScores()
	filepath.Walk(path, compilePath(path, opts, collector))

	PrepareLilypond(opts)
	for i := range scores {
		scores[i].GenerateOutputPath(opts)

		if !helpers.Exists(scores[i].OutputPath) ||
			scores[i].LastModified.After(helpers.LastModified(scores[i].OutputPath)) {
			var (
				msg string
				err error
			)
			switch filepath.Ext(scores[i].Path) {
			case ".ly":
				msg, err = Lilypond(&scores[i])
			case ".pdf":
				err = linkPDF(&scores[i])
			}

			if err != nil {
				log.WithFields(log.Fields{
					"message": msg,
					"error":   err,
					"score":   scores[i],
				}).Warning("score failed to compile")
			}
		} else {
			log.WithFields(log.Fields{
				"score":         scores[i],
				"outputVersion": helpers.LastModified(scores[i].OutputPath),
			}).Debug("skipping compilation")
		}
	}
}

func generateScores() func(string, os.FileInfo) error {
	return func(path string, file os.FileInfo) error {
		var (
			score *settings.Score
			err   error
		)
		switch filepath.Ext(path) {
		case ".ly":
			log.WithFields(log.Fields{"path": path}).Debug("adding lilypond file")
			score, err = settings.FromLy(path)
			if score != nil {
				score.LastModified = file.ModTime()
			}

		case ".json":
			if filepath.Base(path) != "ponder.json" {
				log.WithFields(log.Fields{"path": path}).Debug("adding json file")
				score, err = settings.FromJSON(path)
				score.LastModified = helpers.LastModified(score.Path)
			}

		default:
			log.WithFields(log.Fields{"path": path}).Debug("ignoring file")
		}

		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"path":  path,
			}).Warning("unable to parse score settings, skipping...")
		} else if score != nil {
			scores = append(scores, *score)
		}
		return nil
	}
}

func linkPDF(s *settings.Score) (err error) {
	err = helpers.ExistsOrCreate(filepath.Dir(s.OutputPath))
	if err != nil {
		return err
	}
	if helpers.Exists(s.OutputPath) {
		os.Remove(s.OutputPath)
	}
	err = os.Link(s.Path, s.OutputPath)
	return
}
