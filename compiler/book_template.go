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

	parsePartialTemplate(t.New("Packages"), opts.BookPackagesTempl, packagesTempl)
	parsePartialTemplate(t.New("Title"), opts.BookTitleTempl, titleTempl)
	parsePartialTemplate(t.New("Category"), opts.BookCategoryTempl, categoryTempl)
	parsePartialTemplate(t.New("Score"), opts.BookScoreTempl, scoreTempl)

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

func parsePartialTemplate(t *template.Template, source, fallback string) {
	var err error
	if source != "" {
		_, err = t.Parse(source)
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

const bookTempl = `{{ template "Packages" . }}

{{if ne .Settings.Name ""}}\title{ {{.Settings.Name}} }{{end}}
{{if ne .Settings.Author ""}}\author{ {{.Settings.Author}} }{{end}}
\date{\today}

\begin{document}
{{ template "Title" . }}

{{range $i, $cat := .Categories}}
{{ template "Category" . }}
{{range $.Scores}}{{if in $cat .Categories }}{{template "Score" . }}{{end}}{{end}}
{{end}}

{{if not .Settings.HideUncategorized }}{{ if unknown .Scores }}
{{ if ne .Settings.UncategorizedChapter "" }}{{$title := .Settings.UncategorizedChapter}}{{else}}{{$title := "Others"}}{{ template "Category" $title }}{{end}}
{{range .Scores}}{{ if eq (len .Categories) 0 }}{{template "Score" . }}{{end}}{{end}}
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
