package kjdata

import (
	"html/template"

	"github.com/zach0zhang/kjblog/markdown"
	"github.com/zach0zhang/kjblog/mdfile"
)

type Body struct {
	Title    string
	Keywords string
	Articles mdfile.Articles
}

func (body Body) MarkdownToHTML(input string) template.HTML {
	return markdown.HTML(input)
}
