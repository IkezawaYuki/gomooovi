package main

import (
	"gomooovi/routes"
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", routes.Index)

	server := &http.Server{
		Addr:  "127.0.0.1:8080",
		Handler:mux,
	}

	server.ListenAndServe()
}
