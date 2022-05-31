package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	http.HandleFunc("/", someFunc)
	http.HandleFunc("/hash", hashController)
	http.ListenAndServe(":8080", mux)
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello universe"))
}

func hashController(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content Type", "text/html")

	var hashMe io.ReadCloser = req.Body

	if hashMe != nil {
		var bytes []byte
		_, err := hashMe.Read(bytes)
		if err != nil {
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
				return
			}
		}
	}
}
