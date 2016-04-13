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
	"github.com/alecthomas/template"
	"github.com/jjdekker/ponder/helpers"
	"github.com/jjdekker/ponder/settings"
)

// TODO: Add git version
// TODO: Support multiple authors
// TODO: Support categories
var bookTempl = `
\documentclass[11pt,fleqn]{book}
\usepackage[utf8]{inputenc}
\usepackage{pdfpages}
\usepackage[space]{grffile}
\usepackage{hyperref}

{{if ne .Settings.Name ""}}\\title{ {{.Settings.Name}} }{{end}}
{{if ne .Settings.Author ""}}\\author{ {{.Settings.Author}} }{{end}}
\date{\today}

\begin{document}
\maketitle

{{range $i, $cat := .Categories}}
\chapter{{printf "{"}}{{ . }}{{printf "}"}}
\newpage
{{range $.Scores}}
{{if in $cat .Categories }}
\phantomsection
\addcontentsline{toc}{section}{{printf "{"}}{{ .Name }}{{printf "}"}}
\includepdf[pages=-]{{printf "{"}}{{.OutputPath}}{{printf "}"}}
{{end}}
{{end}}
{{end}}

{{ if unknown .Scores }} \chapter{Others} \newpage {{end}}
{{range .Scores}}
{{ if eq (len .Categories) 0 }}
\phantomsection
\addcontentsline{toc}{section}{{printf "{"}}{{ .Name }}{{printf "}"}}
\includepdf[pages=-]{{printf "{"}}{{.OutputPath}}{{printf "}"}}
{{end}}
{{end}}
\end{document}
`

// MakeBook will combine all scores into a single songbook
// generated using LaTeX.
func MakeBook(path string, opts *settings.Settings) {
	// Everything needs to be compiled
	CompileDir(path, opts)
	// Compile the book template
	var templ = template.Must(template.New("songBook").Funcs(template.FuncMap{
		"in": helpers.InSlice,
		"unknown": unknownCategories,
	}).Parse(bookTempl))

	texPath := filepath.Join(opts.OutputDir, "songbook.tex")
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

	cmd = exec.Command("latexmk", "-c", "-cd", texPath)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.WithFields(log.Fields{
			"message": string(out),
			"error":   err,
		}).Error("failed to clean songbook latex files")
	}
	// err = os.Remove(texPath)
	helpers.Check(err, "could not remove songbook latex template")
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
