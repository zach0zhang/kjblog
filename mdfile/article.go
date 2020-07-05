package mdfile

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"syscall"
	"time"

	"github.com/zach0zhang/kjblog/kjlog"

	"github.com/zach0zhang/kjblog/kjconfig"
)

type Articles []Article

type Article struct {
	// 文章的标题
	Title string

	// 作者姓名
	Author string

	// 作者个人主页
	HomePage string

	// 创建时间
	CreatedAt time.Time

	// 最后更新时间
	UpdatedAt time.Time

	// 所属分类名称
	Category string

	// 文章内容
	Body string

	// 文章在服务器路径
	Path string
}

func getArticlesSpecifiedCategory(category *Category) Articles {
	path := filepath.Join(kjconfig.Cfg.DocPath, category.Path)
	//fmt.Println(path)
	//fmt.Println(kjconfig.Cfg.DocPath)
	filesInfo, err := ioutil.ReadDir(path)
	if os.IsNotExist(err) || len(filesInfo) == 0 {
		fmt.Println("is not exist or filesinfo == 0")
		return nil
	}

	number := 0
	articles := make([]Article, 0)

	for _, info := range filesInfo {
		if info.IsDir() {
			continue
		}

		fileName := info.Name()

		ext := filepath.Ext(fileName)
		if ext != ".md" || fileName == "README.md" {
			continue
		}

		article := getArticleContent(filepath.Join(path, fileName))
		if reflect.DeepEqual(article, Article{}) {
			continue
		}

		article.Category = category.Path
		article.Path = strings.TrimSuffix(info.Name(), ext)

		articles = append(articles, article)
		number++
	}

	category.Number = number
	return articles
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

func getArticleContent(path string) Article {
	article := Article{}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		kjlog.Error("read file error:", path, err)
		return article
	}

	str := string(bytes)
	if str == "" {
		kjlog.Error("file:", path, "is empty")
	}

	article.Body = str

	article.Title = strings.TrimSuffix(filepath.Base(path), ".md")
	// 暂时写死
	article.Author = "zach"
	article.HomePage = "/"

	finfo, _ := os.Stat(path)
	// Sys()返回的是interface{}，所以需要类型断言，不同平台需要的类型不一样，linux上为*syscall.Stat_t
	stat_t := finfo.Sys().(*syscall.Stat_t)
	article.CreatedAt = timespecToTime(stat_t.Ctim)
	article.UpdatedAt = timespecToTime(stat_t.Mtim)

	return article
}

// Len
func (a Articles) Len() int {
	return len(a)
}

// Swap 实现的 sort 接口
func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Articles) Less(i, j int) bool {
	if a[i].UpdatedAt.After(a[j].UpdatedAt) {
		return true
	}

	if a[i].CreatedAt.After(a[j].CreatedAt) {
		return true
	}

	return false
}
