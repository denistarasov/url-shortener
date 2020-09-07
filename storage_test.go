package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStorage_Simple(t *testing.T) {
	s := NewStorage()
	URL := "github.com"
	shortenedURL := s.Post(URL)
	actualURL, exists := s.Get(shortenedURL)
	require.True(t, exists)
	require.Equal(t, URL, actualURL)
}

func TestStorage_GetNonexistentURL(t *testing.T) {
	s := NewStorage()
	shortenedURL := "localhost/abc"
	_, exists := s.Get(shortenedURL)
	require.False(t, exists)
}

func TestStorage_Persistence(t *testing.T) {
	s := NewStorage()
	URL := "github.com"
	shortenedURL := s.Post(URL)
	nextShortenedURL := s.Post(URL)
	require.Equal(t, shortenedURL, nextShortenedURL)
}
