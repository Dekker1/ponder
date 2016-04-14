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

func parseBookTemplate(opts *settings.Settings) (t *template.Template, err error) {
	t = template.New("Songbook")
	t.Funcs(template.FuncMap{
		"in":      helpers.InSlice,
		"unknown": unknownCategories,
	})

	t, err = t.Parse(bookTempl)
	if err != nil {
		log.WithFields(log.Fields{
			"template": t,
			"source":   bookTempl,
			"error":    err,
		}).Fatal("songbook template failed to compile")
	}
	return
}

const bookTempl = `
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
{{range $.Scores}}{{if in $cat .Categories }}
\phantomsection
\addcontentsline{toc}{section}{{printf "{"}}{{ .Name }}{{printf "}"}}
\includepdf[pages=-]{{printf "{"}}{{.OutputPath}}{{printf "}"}}
{{end}}{{end}}{{end}}

{{if not .Settings.HideUncategorized }}{{ if unknown .Scores }}
\chapter{{printf "{"}}{{ if ne .Settings.UncategorizedChapter "" }}{{.Settings.UncategorizedChapter}}{{else}}Others{{end}}{{printf "}"}} \newpage {{end}}
{{range .Scores}}
{{ if eq (len .Categories) 0 }}
\phantomsection
\addcontentsline{toc}{section}{{printf "{"}}{{ .Name }}{{printf "}"}}
\includepdf[pages=-]{{printf "{"}}{{.OutputPath}}{{printf "}"}}
{{end}}{{end}}{{end}}
\end{document}
`
