package mdfile

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type ListMap struct {
	Articles   Articles
	Categories []Category
}

func (list *ListMap) Reload() {
	Model = Listnew()
}

func (list *ListMap) CategoriesAll() Categories {

	return list.Categories
}

func (list *ListMap) ArticlesAll() Articles {
	return list.Articles
}

func (list *ListMap) ArticleByPath(path string) (Article, error) {
	if path != "" {
		for _, article := range list.Articles {
			if article.Path == strings.Trim(path, "/") {
				return article, nil
			}
		}
	}

	return Article{}, errors.New("can not found article")
}

func (list *ListMap) ArticlesByCategory(name string) Articles {
	articles := make(Articles, 0)
	for _, article := range list.Articles {
		if strings.ToLower(article.Category) == strings.ToLower(strings.Trim(name, "/")) {
			articles = append(articles, article)
		}
	}

	return articles
}

func newListMap() *ListMap {
	categoies, err := parseCategories()
	if err != nil {
		//kjlog.Error("parseCategories exec error: ", err)
		fmt.Println(err)
	}
	list := ListMap{
		Categories: categoies,
	}

	list.initArticles()

	return &list
}

func (list *ListMap) initArticles() {
	articles := make(Articles, 0)

	for _, category := range list.Categories {
		//fmt.Println(category.Path)
		article := getArticlesSpecifiedCategory(&category)
		mergeArticles := make(Articles, len(articles)+len(article))
		copy(mergeArticles, articles)
		copy(mergeArticles[len(articles):], article)
		articles = mergeArticles
	}

	// for i, a := range articles {
	// 	fmt.Println(i, a.Title)
	// }

	sort.Sort(&articles)
	list.Articles = articles
}
