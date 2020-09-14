package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStorage_Simple(t *testing.T) {
	s := NewStorage()
	URL := "github.com"
	shortenedURL, err := s.Post(URL, "")
	require.NoError(t, err)
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
	shortenedURL, err := s.Post(URL, "")
	require.NoError(t, err)
	nextShortenedURL, err := s.Post(URL, "")
	require.NoError(t, err)
	require.Equal(t, shortenedURL, nextShortenedURL)
}

func TestStorage_CustomLink(t *testing.T) {
	s := NewStorage()
	URL := "github.com"
	customLink := "github"
	shortenedURL, err := s.Post(URL, customLink)
	require.NoError(t, err)
	require.Equal(t, customLink, shortenedURL)
}

func TestStorage_UsedCustomLink(t *testing.T) {
	s := NewStorage()
	URL := "github.com"
	shortenedURL, err := s.Post(URL, "")
	require.NoError(t, err)
	_, err = s.Post("google.com", shortenedURL)
	require.Error(t, err)
}
