package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func testGetHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		title := r.URL.Path[len("/get/"):]
		p, err := loadPage(title)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte(`Not Found`))
		} else {
			w.Write(p.Body)
		}
	default:
		w.WriteHeader(405)
		w.Write([]byte(`Allow: GET`))
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello, Luan! I'm vivão.`))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`pong`))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello, Luan! This is a test Page.`))
}

func testPostHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		title := r.URL.Path[len("/post/"):]
		file, _, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		p := &Page{Title: title, Body: fileBytes}
		p.save()

		w.Write([]byte(`File uploaded.`))
	default:
		w.WriteHeader(405)
		w.Write([]byte(`Allow: POST`))
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/get/", testGetHandler)
	http.HandleFunc("/post/", testPostHandler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
