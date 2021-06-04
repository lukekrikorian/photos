package main

import (
	"fmt"
	"net/http"
	"photos/files"
	"text/template"
	"time"
)

var (
	tmpls, err = template.ParseFiles("static/index.html")
	index      = template.Must(tmpls, err).Lookup("index")
)

type Injection struct {
	Thumbnails []int
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, Injection{
		Thumbnails: make([]int, files.Count),
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
