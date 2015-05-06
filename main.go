package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
      <html>
      <head>
      <title>Chat</title>
      </head>
      <body>
      Let's chat!
      [ 10 ]Chapter 1
      </body>
      </html>
      `))
	})
	// start the web server
	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
