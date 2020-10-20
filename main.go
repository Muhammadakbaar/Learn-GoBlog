package main

import (
	"net/http"
)

const message = "Hello World"

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf" )
		w.WriteHeader(http.StatusOK)
		w.Write([] byte (message))
	})
	http.ListenAndServe(":8080", mux)

	
}