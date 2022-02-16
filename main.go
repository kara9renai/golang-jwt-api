package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type post struct {
	Title string
	Tag   string
	URL   string
}

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "ふつうのLinuxプログラミング",
		Tag:   "Linux",
		URL:   "https://www.sbcr.jp/product/4797386479/",
	}
	json.NewEncoder(w).Encode(post)
})

func main() {
	r := mux.NewRouter()
	r.Handle("/public", nil)
}
