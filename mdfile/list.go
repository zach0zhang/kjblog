package mdfile

import (
	"encoding/json"
	"os"
)

var Model List

type List interface {
	CategoriesAll() Categories
	ArticlesAll() Articles
	ArticleByPath(string) (Article, error)
	ArticlesByCategory(string) Articles
	Reload()
}

type Categories []Category

type Category struct {
	// 分类名称
	Title string
	// 分类下文章数量
	Number int

	// 分类文件目录
	Path string

	// 分类的描述
	Description string
}

func Listnew() List {
	list := newListMap()

	return list
}

func parseCategories() (Categories, error) {
	jsonData := struct {
		Category []Category
	}{}

	file, err := os.Open("kjconfig/categories.json")
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonData)
	if err != nil {
		return nil, err
	}

	return jsonData.Category, nil
}
