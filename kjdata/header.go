package kjdata

import (
	"github.com/zach0zhang/kjblog/kjconfig"
	"github.com/zach0zhang/kjblog/mdfile"
)

type Header struct {
	WebName    string
	Categories mdfile.Categories
}

var HeaderData Header

func InitHeaderData() {
	HeaderData.WebName = kjconfig.Cfg.WebName
	HeaderData.Categories = mdfile.Model.CategoriesAll()
}
