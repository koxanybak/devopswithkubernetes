package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var createdAt time.Time
var expiresAt time.Time

var filePath string = "assets/daily.jpg"

// ImageHandler gets and responds with the daily image
func ImageHandler(w http.ResponseWriter, r *http.Request) {
	// Check if a day has passed
	if createdAt.IsZero() || expiresAt.Before(time.Now()) {
		res, err := http.Get("https://picsum.photos/300")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer res.Body.Close()

		createdAt = time.Now()
		expiresAt = createdAt.Add(time.Hour*24)

		// Cache the image
		file, err := os.Create(filePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		io.Copy(file, res.Body)
	}

	// Get the image from cache
	daily, err := os.Open(filePath)
	info, err := daily.Stat()
	size := info.Size()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond
	w.Header().Set("Content-Length", fmt.Sprint(size))
	w.Header().Set("Content-Type", "image/jpeg")
	io.Copy(w, daily)
}