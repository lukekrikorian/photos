package main

import (
	"fmt"
	"net/http"
	"photos/config"
	"photos/files"
	"text/template"
	"time"
)

var (
	tfiles, err = template.ParseFiles("static/index.html", "static/feed.xml")
	tmpls       = template.Must(tfiles, err)
	index       = tmpls.Lookup("index")
	feed        = tmpls.Lookup("feed")
)

type RSS struct {
	Author  string
	BaseURL string
	Dates   []string
}

type Injection struct {
	Thumbnails []int
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, Injection{
		Thumbnails: make([]int, files.Count),
	})
}

func serveFeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/rss+xml")
	feed.Execute(w, RSS{
		Author:  config.Config.Author,
		BaseURL: config.Config.BaseURL,
		Dates:   files.GetFileDates("static/photos"),
	})
}

func handleUploads(w http.ResponseWriter, r *http.Request) {
	defer http.ServeFile(w, r, "static/upload.html")

	if r.Method == http.MethodGet {
		return
	}

	r.ParseMultipartForm(10 << 20) // 10 MB
	file, header, err := r.FormFile("photo")

	if err != nil {
		return
	}
	defer file.Close()

	files.SavePhoto(&file, header)
}

func main() {
	fmt.Println("Starting!")

	handler := http.DefaultServeMux

	static := http.FileServer(http.Dir("./static"))
	handler.Handle("/static/", http.StripPrefix("/static/", static))

	handler.HandleFunc("/", serveIndex)
	handler.HandleFunc("/upload", handleUploads)
	handler.HandleFunc("/feed", serveFeed)

	s := &http.Server{
		Handler:      handler,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	if err = s.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
