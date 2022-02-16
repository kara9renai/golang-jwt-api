package main

import (
	"encoding/json"
	"log"
	"net/http"

	"gihub.com/kara9renai/golang-jwt-api/auth"
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

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title: "VGolangとGoogle Cloud Vision APIで画像から文字認識するCLIを速攻でつくる",
		Tag:   "Go",
		URL:   "https://qiita.com/po3rin/items/bf439424e38757c1e69b",
	}
	json.NewEncoder(w).Encode(post)
})

func main() {
	const port = ":8080"

	r := mux.NewRouter()
	r.Handle("/public", public)
	r.Handle("/private", auth.JwtMiddleware.Handler(private))
	r.Handle("/auth", auth.GetTokenHandler)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
