package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Page is the webpage
type Page struct {
	Title string
	Body  []byte
}

var (
	promRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "debugtools_requests",
			Help: "The total number of request events",
		},
		[]string{"header_host", "route", "client"},
	)
	promFiles = promauto.NewCounter(prometheus.CounterOpts{
		Name: "debugtools_files_uploaded",
		Help: "The total number of files uploaded",
	})
	promErrors = promauto.NewCounter(prometheus.CounterOpts{
		Name: "debugtools_errors",
		Help: "The total number of errors",
	})
)

func monitorRoute(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			promErrors.Inc()
			return
		}

		userIP := net.ParseIP(ip)
		if userIP == nil {
			promErrors.Inc()
			return
		}

		promRequests.With(prometheus.Labels{
			"header_host": r.Host,
			"route":       r.URL.Path,
			"client":      userIP.String(),
		}).Inc()
		f(w, r)
	}
}

func (p *Page) save() error {
	filename := p.Title + ".save"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".save"
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
			promErrors.Inc()
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
			promErrors.Inc()
			panic(err)
		}
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		p := &Page{Title: title, Body: fileBytes}
		p.save()

		promFiles.Inc()
		w.Write([]byte(`File uploaded.`))
	default:
		w.WriteHeader(405)
		w.Write([]byte(`Allow: POST`))
	}
}

func main() {
	http.HandleFunc("/health", monitorRoute(healthHandler))
	http.HandleFunc("/ping", monitorRoute(pingHandler))
	http.HandleFunc("/", monitorRoute(homeHandler))
	http.HandleFunc("/get/", monitorRoute(testGetHandler))
	http.HandleFunc("/post/", monitorRoute(testPostHandler))
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":5000", nil))
}
