package main

import (
	"embed"
	_ "embed"
	"log"
	"net/http"
)

//go:embed stylesheet.css
var staticFiles embed.FS

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO return 404 if the wrong url is queried
		articles := generateArticles(40)
		document := indexDocument(articles)	
		_, err := w.Write(document.Render())
		if err != nil {
			// What should the error handling be here?
			log.Fatal(err)
		}
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFiles))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
