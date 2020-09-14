package main

import (
	"fmt"
	"math/rand"
)

const (
	shortURLLength = 6
)

var (
	symbols                      = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	CustomLinkAlreadyExistsError = fmt.Errorf("short link already exists")
)

type Storage struct {
	shortToFullURLs map[string]string
	fullURLsToShort map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		shortToFullURLs: make(map[string]string),
		fullURLsToShort: make(map[string]string),
	}
}

func generateRandomStr(length int) string {
	buf := make([]rune, length)
	for i := range buf {
		buf[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(buf)
}

func (s *Storage) Get(shortURL string) (string, bool) {
	fullURL, exists := s.shortToFullURLs[shortURL]
	return fullURL, exists
}

func (s *Storage) Post(fullURL, customLink string) (string, error) {
	shortStr, exists := s.fullURLsToShort[fullURL]
	if exists {
		return shortStr, nil
	}
	if customLink != "" {
		if _, exists := s.shortToFullURLs[customLink]; exists {
			return "", CustomLinkAlreadyExistsError
		}
		shortStr = customLink
	} else {
		for {
			shortStr = generateRandomStr(shortURLLength)
			if _, exists := s.shortToFullURLs[shortStr]; !exists {
				break
			}
		}
	}
	s.shortToFullURLs[shortStr] = fullURL
	s.fullURLsToShort[fullURL] = shortStr
	return shortStr, nil
}
