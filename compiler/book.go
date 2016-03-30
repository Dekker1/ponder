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

	log "github.com/Sirupsen/logrus"
	"github.com/alecthomas/template"
	"github.com/jjdekker/ponder/helpers"
	"github.com/jjdekker/ponder/settings"
)

// TODO: Add git version
// TODO: Support multiple authors
// TODO: Support categories
// TODO: Add working TOC
var bookTempl = `
\documentclass[a4paper,11pt]{article}
\usepackage[utf8]{inputenc}
\usepackage{pdfpages}
\usepackage[space]{grffile}

{{if ne .Settings.Name ""}}\\title{ {{.Settings.Name}} }{{end}}
{{if ne .Settings.Author ""}}\\author{ {{.Settings.Author}} }{{end}}
\date{\today}

\begin{document}
\maketitle

{{range .Scores}}
\includepdf[pages=-]{{printf "{"}}{{call $.OutputPath .Path $.Settings}}{{printf "}"}}
{{end}}

\end{document}
`

// MakeBook will combine all scores into a single songbook
// generated using LaTeX.
func MakeBook(path string, opts *settings.Settings) {
	// Everything needs to be compiled
	CompileDir(path, opts)
	// Compile the book template
	var templ = template.Must(template.New("songBook").Parse(bookTempl))

	texPath := filepath.Join(opts.OutputDir, "songbook.tex")
	log.WithFields(log.Fields{
		"path": texPath,
	}).Info("compiling songbook template")
	f, err := os.Create(texPath)
	helpers.Check(err, "could not create songbook texfile")
	err = templ.Execute(f, &struct {
		Scores     []*settings.Score
		Settings   *settings.Settings
		OutputPath func(string, *settings.Settings) string
	}{
		Scores:     scores,
		Settings:   opts,
		OutputPath: outputPath,
	})
	helpers.Check(err, "error executing book template")
	f.Close()

	// cmd := exec.Command("pdflatex", "-output-directory="+opts.OutputDir, texPath)
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
	err = os.Remove(texPath)
	helpers.Check(err, "could not remove songbook latex template")
}
