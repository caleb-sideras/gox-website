package data

import (
	"html/template"

	"github.com/caleb-sideras/gox/.gox/render"
	"github.com/caleb-sideras/gox/src/global"
	"github.com/caleb-sideras/gox/src/pages/docs"
)

func Render() render.DynamicT {

	tmpls := []string{"templates/components/nav.html"}

	markdownDocsContent := docs.DocsData{
		ActiveTabId:  "docs",
		ActiveDocsId: "data",
	}

	tmpl := template.Must(global.MarkdownToHTML("pages/docs/_markdown/data.md"))

	return render.DynamicT{tmpls, markdownDocsContent, tmpl}
}
