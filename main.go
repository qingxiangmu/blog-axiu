package main

import (
	"html/template"
	"net/http"
	"os"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "axiu's blog"
	indexData.Desc = "now is introduction"
	t := template.New("index.html")
	path, _ := os.Getwd()
	home := path + "/template/home.html"
	header := path + "/template/home.html"
	footer := path + "/template/layout/footer.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	pagination := path + "/template/layout/pagination.html"
	t, _ = t.ParseFiles(path+"/template/index.html", header, home, footer, personal, post, pagination)
	t.Execute(w, indexData)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9898",
	}
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
	}
}
