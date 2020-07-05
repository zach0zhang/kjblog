package controller

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/zach0zhang/kjblog/kjdata"
	"github.com/zach0zhang/kjblog/mdfile"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("resources/pages/index.html",
		"resources/pages/layouts/_head.html",
		"resources/pages/layouts/_header.html"))
	htmlData := kjdata.KjData{
		HeaderData: kjdata.HeaderData,
		BodyData: kjdata.Body{
			Title:    "博客列表",
			Keywords: "博客列表",
			Articles: mdfile.Model.ArticlesAll(),
		},
	}
	t.Execute(w, htmlData)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("resources/pages/notfound.html",
		"resources/pages/layouts/_head.html",
		"resources/pages/layouts/_header.html"))
	htmlData := kjdata.KjData{
		HeaderData: kjdata.HeaderData,
	}
	t.Execute(w, htmlData)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	categoryName := r.FormValue("Name")
	if categoryName == "/" {
		http.Redirect(w, r, "/", 302)
	}

	articles := mdfile.Model.ArticlesByCategory(categoryName)

	if len(articles) > 0 {
		t := template.Must(template.ParseFiles("resources/pages/index.html",
			"resources/pages/layouts/_head.html",
			"resources/pages/layouts/_header.html"))
		htmlData := kjdata.KjData{
			HeaderData: kjdata.HeaderData,
			BodyData: kjdata.Body{
				Title:    "分类 | " + categoryName,
				Keywords: categoryName,
				Articles: articles,
			},
		}
		t.Execute(w, htmlData)
	} else {
		NotFound(w, r)
		//http.Redirect(w, r, "/notfound", 302)
	}
}

func ToKeywords(works ...string) string {
	return strings.Join(works, ",")
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("Path")

	article, err := mdfile.Model.ArticleByPath(path)
	if err != nil {
		NotFound(w, r)
		return
	}

	var articles mdfile.Articles
	articles = append(articles, article)

	t := template.Must(template.ParseFiles("resources/pages/article.html",
		"resources/pages/layouts/_head.html",
		"resources/pages/layouts/_header.html"))

	htmlData := kjdata.KjData{
		HeaderData: kjdata.HeaderData,
		BodyData: kjdata.Body{
			Title:    article.Title,
			Keywords: ToKeywords(article.Title, article.Path),
			Articles: articles,
		},
	}
	t.Execute(w, htmlData)

}
