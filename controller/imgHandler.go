package controller

import (
	"io"
	"net/http"
	"os"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	returnImage(w, path)
}

func returnImage(w http.ResponseWriter, path string) {
	file, err := os.Open(path)
	if err != nil {
		//noImage(w)
		return
	}
	io.Copy(w, file)
	file.Close()
	return
}
