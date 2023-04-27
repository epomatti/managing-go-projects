package main

import (
	"embed"
	"fmt"
	"io/fs"
	"main/numbers"
	"net/http"
)

//go:embed html/index.html
var indexHTML embed.FS

//go:embed css/styles.css
var stylesCSS embed.FS

func handler(fsys embed.FS, subfolder string) http.Handler {
	html, _ := fs.Sub(fs.FS(fsys), subfolder)
	return http.FileServer(http.FS(html))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/text")
	sum := numbers.Sum()
	result := fmt.Sprintf("Success: %v", sum)
	w.Write([]byte(result))
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", handler(indexHTML, "html"))
	mux.Handle("/styles.css", handler(stylesCSS, "css"))
	mux.HandleFunc("/numbers", handleRequest)
	http.ListenAndServe(":5000", mux)
}
