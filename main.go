package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	URL       string `json:"url"`
	CustomURL string `json:"custom_url"`
}

type Response struct {
	URL string `json:"url"`
}

func main() {
	r := mux.NewRouter()
	s := NewStorage()
	r.HandleFunc("/shorten", s.ShorteningURLHandler).Methods("POST")
	r.HandleFunc("/{key}", s.RedirectionHandler).Methods("GET")
	srv := &http.Server{
		Addr:    ":8100",
		Handler: r,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Storage) ShorteningURLHandler(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	request := Request{}
	err := d.Decode(&request)
	if err != nil {
		log.Fatalf("decode error: %v", err)
	}
	_, err = url.ParseRequestURI(request.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte("URL is not valid"))
		if err != nil {
			log.Fatalf("write error: %v", err)
		}
		return
	}
	var response Response
	response.URL, err = s.Post(request.URL, request.CustomURL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			log.Fatalf("write error: %v", err)
		}
		return
	}
	data, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("encode error: %v", err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("write error: %v", err)
	}
}

func (s *Storage) RedirectionHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/")
	URL, exists := s.Get(key)
	if !exists {
		log.Infof("client requested non-existent key: %s", key)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	http.Redirect(w, r, URL, http.StatusSeeOther)
}
