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
	"path/filepath"
	"text/template"

	log "github.com/Sirupsen/logrus"
	"github.com/jjdekker/ponder/helpers"
	"github.com/jjdekker/ponder/settings"
)

// parseBookTemplate parses all partial templates for the book
func parseBookTemplate(opts *settings.Settings) (t *template.Template, err error) {
	t = template.New("Songbook")
	t.Funcs(template.FuncMap{
		"in":      helpers.InSlice,
		"unknown": unknownCategories,
	})

	parsePartialTemplate(t.New("packages.tex.tmpl"),
		filepath.Join(opts.BookTemplateDir, "packages.tex.tmpl"), packagesTempl)
	parsePartialTemplate(t.New("title.tex.tmpl"),
		filepath.Join(opts.BookTemplateDir, "title.tex.tmpl"), titleTempl)
	parsePartialTemplate(t.New("category.tex.tmpl"),
		filepath.Join(opts.BookTemplateDir, "category.tex.tmpl"), categoryTempl)
	parsePartialTemplate(t.New("score.tex.tmpl"),
		filepath.Join(opts.BookTemplateDir, "score.tex.tmpl"), scoreTempl)

	_, err = t.Parse(bookTempl)
	if err != nil {
		log.WithFields(log.Fields{
			"template": t,
			"source":   bookTempl,
			"error":    err,
		}).Fatal("songbook template failed to parse")
	}
	return
}

func parsePartialTemplate(t *template.Template, sourceFile, fallback string) {
	var err error
	if helpers.Exists(sourceFile) {
		_, err = t.ParseFiles(sourceFile)
	} else {
		_, err = t.Parse(fallback)
	}
	if err != nil {
		log.WithFields(log.Fields{
			"source": packagesTempl,
			"error":  err,
		}).Fatal("packages partial template failed to parse")
	}
}

const bookTempl = `{{ template "packages.tex.tmpl" . }}

{{if ne .Settings.Name ""}}\title{ {{.Settings.Name}} }{{end}}
{{if ne .Settings.Author ""}}\author{ {{.Settings.Author}} }{{end}}
\date{\today}

\begin{document}
{{ template "title.tex.tmpl" . }}

{{range $i, $cat := .Categories}}
{{ template "category.tex.tmpl" . }}
{{range $.Scores}}{{if in $cat .Categories }}{{template "score.tex.tmpl" . }}{{end}}{{end}}
{{end}}

{{if not .Settings.HideUncategorized }}{{ if unknown .Scores }}
{{ if ne .Settings.UncategorizedChapter "" }}{{$title := .Settings.UncategorizedChapter}}{{else}}{{$title := "Others"}}{{ template "category.tex.tmpl" $title }}{{end}}
{{range .Scores}}{{ if eq (len .Categories) 0 }}{{template "score.tex.tmpl" . }}{{end}}{{end}}
{{end}}{{end}}
\end{document}
`

const packagesTempl = `\documentclass[11pt,fleqn]{book}
\usepackage[utf8]{inputenc}
\usepackage{pdfpages}
\usepackage[space]{grffile}
\usepackage{hyperref}`

const titleTempl = `\maketitle`

const categoryTempl = `\chapter{{printf "{"}}{{ . }}{{printf "}"}}\newpage`

const scoreTempl = `\includepdf[addtotoc={1,section,1,{{ printf "{%s}" .Name }},}, pages=-]{{printf "{%s}" .OutputPath}}`
