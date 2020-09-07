package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type Message struct {
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
	req := Message{}
	err := d.Decode(&req)
	if err != nil {
		log.Fatalf("decode error: %v", err)
	}
	req.URL = s.Post(req.URL)
	data, err := json.Marshal(req)
	if err != nil {
		log.Fatalf("encode error: %v", err)
	}
	_, err = w.Write(data)
	if err != nil {
		log.Fatalf("write error: %v", err)
	}
}

func (s *Storage) RedirectionHandler(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.RequestURI, "/")
	URL, exists := s.Get(key)
	if !exists {
		log.Infof("client requested non-existent key: %s", key)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	http.Redirect(w, r, URL, http.StatusSeeOther)
}
