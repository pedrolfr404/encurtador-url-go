package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var (
	urlStore = make(map[string]string)
	mu       sync.Mutex
)

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	originalUrl := r.URL.Query().Get("url")

	if originalUrl == "" {
		http.Error(w, "Missing 'url' parameter", http.StatusBadRequest)
		return
	}

	if !(strings.HasPrefix(originalUrl, "https://")) || !(strings.HasPrefix(originalUrl, "http://")) {
		originalUrl = "http://" + originalUrl
	}
	fmt.Printf(originalUrl)
	encryptedUrl := encrypt(originalUrl)
	shortId := generateShortId()

	mu.Lock()
	urlStore[encryptedUrl] = originalUrl
	mu.Unlock()

	shortUrl := fmt.Sprintf("https://localhost:8080/%s", shortId)
	fmt.Fprintf(w, "A URL encurtada desta url original é: %s", shortUrl)
}
