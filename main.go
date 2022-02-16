package main

import (
	"encoding/json"
	"log"
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
	const port = ":8080"

	r := mux.NewRouter()
	r.Handle("/public", public)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
