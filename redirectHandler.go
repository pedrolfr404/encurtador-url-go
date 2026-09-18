package main

import "net/http"

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortId := r.URL.Path[1:]

	mu.Lock()
	encryptedUrl, ok := urlStore[shortId]
	mu.Unlock()

	if !ok {
		http.Error(w, "URL inexistente", http.StatusNotFound)
		return
	}

	decryptedUrl := decrypt(encryptedUrl)
	http.Redirect(w, r, decryptedUrl, http.StatusFound)
}
