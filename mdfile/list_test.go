package mdfile

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("test arseCategories", testparseCategories)
	t.Run("test getArticleContent", testgetArticleContent)
}

func testparseCategories(t *testing.T) {
	Categories, err := parseCategories()
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("show:")
	for _, value := range Categories {
		fmt.Println(value)
	}
}

func testgetArticleContent(t *testing.T) {
	article := getArticleContent("/home/zach/go/src/github.com/zach0zhang/kjblog/resources/blog-docs/algorithm/array-diagonal-traverse.md")
	fmt.Println(article.Author)
	fmt.Println(article.Title)
	fmt.Println(article.CreatedAt, article.UpdatedAt)
}
