package main

import (
	"gomooovi/routes"
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", routes.Index)
	mux.HandleFunc("/products/search", routes.ProductSearch)

	server := &http.Server{
		Addr:  config.Address,
		Handler:mux,
	}

	server.ListenAndServe()
}
