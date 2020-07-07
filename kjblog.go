package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/zach0zhang/kjblog/controller"
	"github.com/zach0zhang/kjblog/kjdata"
	"github.com/zach0zhang/kjblog/kjlog"
	"github.com/zach0zhang/kjblog/mdfile"

	"github.com/zach0zhang/kjblog/kjconfig"
)

func version() string {
	return "0.1"
}

func init() {
	var cmdConfig kjconfig.Configuration
	var printVer bool
	var configFile string

	dir, _ := os.Getwd()

	flag.BoolVar(&printVer, "version", false, "print version")
	flag.StringVar(&configFile, "c", "kjconfig/config.json", "specify config path")
	flag.StringVar(&cmdConfig.Address, "a", "0.0.0.0:8080", "specify http server address")
	flag.Int64Var(&cmdConfig.ReadTimeout, "r", 10, "read timeout")
	flag.Int64Var(&cmdConfig.WriteTimeout, "w", 600, "write timeout")
	flag.StringVar(&cmdConfig.DocPath, "s", dir+"/resources/blog-docs", "blog doc file path")
	flag.StringVar(&cmdConfig.LogFile, "l", "kjlogLog.txt", "log file path")
	flag.IntVar(&cmdConfig.LogOutputLevel, "logLevel", 0, "output level, 0: debug, 1: info, 2: warning. 3: error")
	flag.Parse()
	if printVer {
		fmt.Println(version())
		os.Exit(0)
	}

	err := kjconfig.InitConfig(configFile)
	if err != nil {
		fmt.Println("IninConfig error: ", err)
		fmt.Println("Use the input configuration", cmdConfig)
		fmt.Println("exec [kjblog --help] can see help message")
		kjconfig.Cfg = cmdConfig
	} else {
		err = kjlog.InitLog(kjconfig.Cfg.LogFile, kjconfig.Cfg.LogOutputLevel)
		if err != nil {
			fmt.Println("InitConfig error: ", err)
			os.Exit(1)
		}
	}

	mdfile.Model = mdfile.Listnew()

	kjlog.Debug(kjconfig.Cfg)
}

func main() {
	fmt.Println("kjbolg", version(), "started at", kjconfig.Cfg.Address)

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("resources/static"))))
	//mux.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("resources/pages"))))

	mux.HandleFunc("/", controller.Index)
	mux.HandleFunc("/notfound", controller.NotFound)
	mux.HandleFunc("/category", controller.GetCategory)
	mux.HandleFunc("/article", controller.GetArticle)

	server := &http.Server{
		Addr:           kjconfig.Cfg.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(kjconfig.Cfg.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(kjconfig.Cfg.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	kjdata.InitHeaderData()

	server.ListenAndServe()
}
