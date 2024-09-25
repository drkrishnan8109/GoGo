package main

//Reference: https://medium.com/bootdotdev/create-a-golang-video-streaming-server-using-hls-a-tutorial-f8c7d4545a0f

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// configure the songs directory name and port
	const songsDir = "/Users/dhanyakrishnan/Documents/workspace/hlsstreaming"
	const port = 8080

	// add a handler for the song files
	http.Handle("/", addHeaders(http.FileServer(http.Dir(songsDir))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", songsDir, port)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
