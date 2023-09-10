package main

import (
	"compress/gzip"
	"embed"
	_ "embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed stylesheet.css
var staticFiles embed.FS

//go:embed template.html
var indexTemplate string

func main() {
	articles := generateArticles(40)
	tmpl, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/sanity", func(w http.ResponseWriter, r *http.Request) {
		view := indexDocument(articles)
		_, err := w.Write(view.Render())
		if err != nil {
			log.Fatal(err)
		}
	})
	http.HandleFunc("/sanity-gzip", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		gw := gzip.NewWriter(w)
		defer gw.Close()
		view := indexDocument(articles)
		_, err := gw.Write(view.Render())
		if err != nil {
			log.Fatal(err)
		}
	})
	staticBytes := indexDocument(articles).Render()
	http.HandleFunc("/bytes", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(staticBytes)
		if err != nil {
			log.Fatal(err)
		}
	})
	http.HandleFunc("/bytes-gzip", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		gw := gzip.NewWriter(w)
		defer gw.Close()
		_, err := gw.Write(staticBytes)
		if err != nil {
			log.Fatal(err)
		}
	})
	http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, articles)
		if err != nil {
			log.Fatal(err)
		}
	})
	http.HandleFunc("/template-gzip", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Encoding", "gzip")
		gw := gzip.NewWriter(w)
		defer gw.Close()
		err := tmpl.Execute(gw, articles)
		if err != nil {
			log.Fatal(err)
		}
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
