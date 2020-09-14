package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlers_Simple(t *testing.T) {
	s := NewStorage()

	body := strings.NewReader(`{"url": "http://github.com"}`)
	r, err := http.NewRequest("POST", "http://localhost:8100", body)
	require.NoError(t, err)
	w := httptest.NewRecorder()
	s.ShorteningURLHandler(w, r)
	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	response := Response{}
	err = json.Unmarshal(respBody, &response)
	require.NoError(t, err)

	URL := fmt.Sprintf("http://localhost:8100/%s", response.URL)
	r = httptest.NewRequest("GET", URL, nil)
	w = httptest.NewRecorder()
	s.RedirectionHandler(w, r)
	require.Equal(t, http.StatusSeeOther, w.Code)
}

func TestHandlers_CustomURL(t *testing.T) {
	s := NewStorage()

	body := strings.NewReader(`{"url": "http://github.com", "custom_url": "abc"}`)
	r, err := http.NewRequest("POST", "http://localhost:8100", body)
	require.NoError(t, err)
	w := httptest.NewRecorder()
	s.ShorteningURLHandler(w, r)
	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	response := Response{}
	err = json.Unmarshal(respBody, &response)
	require.NoError(t, err)
	require.Equal(t, response.URL, "abc")

	URL := "http://localhost:8100/abc"
	r = httptest.NewRequest("GET", URL, nil)
	w = httptest.NewRecorder()
	s.RedirectionHandler(w, r)
	require.Equal(t, http.StatusSeeOther, w.Code)
}

func TestHandlers_InvalidURL(t *testing.T) {
	s := NewStorage()

	body := strings.NewReader(`{"url": "github"}`)
	r, err := http.NewRequest("POST", "http://localhost:8100", body)
	require.NoError(t, err)
	w := httptest.NewRecorder()
	s.ShorteningURLHandler(w, r)
	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	require.Equal(t, "URL is not valid", string(respBody))
}

func TestHandlers_CustomLinkAlreadyExists(t *testing.T) {
	s := NewStorage()

	body := strings.NewReader(`{"url": "http://github.com", "custom_url": "abc"}`)
	r, err := http.NewRequest("POST", "http://localhost:8100", body)
	require.NoError(t, err)
	w := httptest.NewRecorder()
	s.ShorteningURLHandler(w, r)
	resp := w.Result()
	respBody, err := ioutil.ReadAll(resp.Body)
	response := Response{}
	err = json.Unmarshal(respBody, &response)
	require.NoError(t, err)
	require.Equal(t, response.URL, "abc")

	body = strings.NewReader(`{"url": "http://google.com", "custom_url": "abc"}`)
	r, err = http.NewRequest("POST", "http://localhost:8100", body)
	require.NoError(t, err)
	w = httptest.NewRecorder()
	s.ShorteningURLHandler(w, r)
	resp = w.Result()
	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
}


func TestHandlers_NonExistentKey(t *testing.T) {
	s := NewStorage()

	URL := "http://localhost:8100/abc"
	r := httptest.NewRequest("GET", URL, nil)
	w := httptest.NewRecorder()
	s.RedirectionHandler(w, r)
	require.Equal(t, http.StatusNotFound, w.Code)
}
