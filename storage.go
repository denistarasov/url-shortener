package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

const (
	shortURLLength = 6
)

var (
	symbols                      = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	CustomLinkAlreadyExistsError = fmt.Errorf("short link already exists")
	DSN                          = "root:1234@tcp(db:3306)/golang?charset=utf8"
	databasePingRetryCount       = 5
)

func initDB() *sql.DB {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i != databasePingRetryCount; i++ {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(time.Second)
		log.Info("couldn't connect to database, retrying...")
	}
	if err != nil {
		log.Fatalf("connection to database wasn't established after %d retries", databasePingRetryCount)
	}

	queries := []string{
		`DROP TABLE IF EXISTS urls;`,

		`CREATE TABLE urls (
		  full_url varchar(255) NOT NULL,
		  short_link varchar(255) NOT NULL,
		  PRIMARY KEY (short_link)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
	}

	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			log.Fatal(err)
		}
	}

	return db
}

type Storage struct {
	db *sql.DB
}

func NewStorage() *Storage {
	return &Storage{
		db: initDB(),
	}
}

func generateRandomStr(length int) string {
	buf := make([]rune, length)
	for i := range buf {
		buf[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(buf)
}

func (s *Storage) findFullURL(shortLink string) (string, bool) {
	row := s.db.QueryRow("SELECT full_url FROM urls WHERE short_link = ?", shortLink)

	var fullURL string
	err := row.Scan(&fullURL)
	if errors.Is(err, sql.ErrNoRows) {
		return "", false
	}
	if err != nil {
		log.Fatal(err)
	}

	return fullURL, true
}

func (s *Storage) findShortLink(fullURL string) (string, bool) {
	row := s.db.QueryRow("SELECT short_link FROM urls WHERE full_url = ?", fullURL)

	var shortLink string
	err := row.Scan(&shortLink)
	if errors.Is(err, sql.ErrNoRows) {
		return "", false
	}
	if err != nil {
		log.Fatal(err)
	}

	return shortLink, true
}

func (s *Storage) insert(fullURL, shortLink string) {
	_, err := s.db.Exec(
		"INSERT INTO urls (full_url, short_link) VALUES (?, ?)",
		fullURL,
		shortLink,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Storage) Get(shortLink string) (string, bool) {
	fullURL, exists := s.findFullURL(shortLink)
	return fullURL, exists
}

func (s *Storage) Post(fullURL, customLink string) (string, error) {
	shortLink, exists := s.findShortLink(fullURL)
	if exists {
		return shortLink, nil
	}
	if customLink != "" {
		if _, exists := s.findFullURL(customLink); exists {
			return "", CustomLinkAlreadyExistsError
		}
		shortLink = customLink
	} else {
		for {
			shortLink = generateRandomStr(shortURLLength)
			if _, exists := s.findFullURL(shortLink); !exists {
				break
			}
		}
	}
	s.insert(fullURL, shortLink)
	return shortLink, nil
}
