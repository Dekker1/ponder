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
	"path/filepath"
	"sort"

	log "github.com/Sirupsen/logrus"
	"github.com/jjdekker/ponder/helpers"
	"github.com/jjdekker/ponder/settings"
)

// MakeBook will combine all scores into a single songbook
// generated using LaTeX.
func MakeBook(path string, opts *settings.Settings) {
	// Everything needs to be compiled
	CompileDir(path, opts)
	// Sort scores
	sort.Sort(settings.ScoresByName{scores})

	getBookResources(opts)

	templ, err := parseBookTemplate(opts)

	texPath := filepath.Join(opts.OutputDir, opts.Name+".tex")
	log.WithFields(log.Fields{
		"path": texPath,
	}).Info("compiling songbook template")
	f, err := os.Create(texPath)
	helpers.Check(err, "could not create songbook texfile")
	err = templ.Execute(f, &struct {
		Scores     *[]settings.Score
		Settings   *settings.Settings
		Categories []string
	}{
		Scores:     &scores,
		Settings:   opts,
		Categories: scoreCategories(&scores),
	})
	helpers.Check(err, "error executing book template")
	f.Close()

	cmd := exec.Command("latexmk", "-silent", "-pdf", "-cd", texPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.WithFields(log.Fields{
			"message": string(out),
			"error":   err,
		}).Fatal("songbook failed to compile")
	}

	cleanBookResources(texPath, opts)
}

// scoreCategories returns a sorted slice of all categories used
// in the given slice of scores
func scoreCategories(scores *[]settings.Score) []string {
	catMap := make(map[string]struct{})
	for i := range *scores {
		for _, cat := range (*scores)[i].Categories {
			catMap[cat] = struct{}{}
		}
	}
	categories := make([]string, 0, len(catMap))
	for cat := range catMap {
		categories = append(categories, cat)
	}
	sort.Strings(categories)
	return categories
}

// unknownCategories returns true if the slice contains any scores with
// unknown categories
func unknownCategories(scores *[]settings.Score) bool {
	for i := range *scores {
		if len((*scores)[i].Categories) == 0 {
			return true
		}
	}
	return false
}

// getBookResources copies the LaTeX resources to the compile directory
func getBookResources(opts *settings.Settings) {
	for i := range opts.LatexResources {
		log.WithFields(log.Fields{"path": opts.LatexResources[i]}).Debug("copying latex resource")

		err := os.Link(opts.LatexResources[i], filepath.Join(opts.OutputDir, filepath.Base(opts.LatexResources[i])))

		if err != nil {
			log.WithError(err).Warning("could not link LaTeX resource")
		}
	}
}

// cleanBookResources removes both LaTeX generated files and copied resources
func cleanBookResources(bookpath string, opts *settings.Settings) {
	var err error
	log.Debug("removing LaTeX resources")
	for i := range opts.LatexResources {
		err = os.RemoveAll(filepath.Join(opts.OutputDir, filepath.Base(opts.LatexResources[i])))
		if err != nil {
			helpers.Check(err, "could not remove latex resource")
		}
	}

	log.Debug("removing LaTeX generated files")
	cmd := exec.Command("latexmk", "-c", "-cd", bookpath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.WithFields(log.Fields{
			"message": string(out),
			"error":   err,
		}).Error("failed to clean songbook latex files")
	}

	if !opts.KeepBookTemplate {
		log.Debug("removing LaTeX template")
		err = os.Remove(bookpath)
	}
	helpers.Check(err, "could not remove songbook latex template")
}
