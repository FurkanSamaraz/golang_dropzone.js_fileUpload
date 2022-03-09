package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/upload", Upload)
	http.ListenAndServe(":8080", nil)
}
func Index(w http.ResponseWriter, r *http.Request) {
	view, _ := template.ParseFiles("index.html")
	view.Execute(w, nil)
}
func Upload(w http.ResponseWriter, r *http.Request) {
	file, header, _ := r.FormFile("file")
	f, _ := os.OpenFile(header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)

}
