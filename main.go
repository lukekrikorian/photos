package main

import (
	"fmt"
	_ "image"
	_ "image/png"
	"log"
	"net/http"
	"photos/files"
	"text/template"
	"time"
)

var (
	index = template.Must(template.ParseFiles("static/index.html")).Lookup("index")
)

type Injection struct {
	Thumbnails []string
	Photos     []string
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	index.Execute(w, Injection{
		Photos:     files.GetFiles("static/photos"),
		Thumbnails: files.GetFiles("static/thumbnails"),
	})
}

func handleUploads(w http.ResponseWriter, r *http.Request) {
	defer http.Redirect(w, r, "/static/upload.html", http.StatusSeeOther)

	if r.Method == http.MethodGet {
		return
	}

	r.ParseForm()

	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("photo")

	if err != nil {
		return
	}
	defer file.Close()

	rotate := false
	if r.Form["rotate"] != nil {
		rotate = true
	}

	files.SavePhoto(&file, header, rotate)
}

func main() {
	files.GenerateThumbnails()
	fmt.Println("Starting!")

	handler := http.DefaultServeMux

	static := http.FileServer(http.Dir("./static"))
	handler.Handle("/static/", http.StripPrefix("/static/", static))

	handler.HandleFunc("/", serveIndex)
	handler.HandleFunc("/photos", handleUploads)

	s := &http.Server{
		Handler:      handler,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
