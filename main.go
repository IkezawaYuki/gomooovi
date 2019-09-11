package main

import (
	"gomooovi/routes"
	"net/http"

	_ "gomooovi/models"
)

func main(){

	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", routes.Index)

	mux.HandleFunc("/login", routes.Login)
	mux.HandleFunc("/signup", routes.Signup)

	mux.HandleFunc("/products/search", routes.ProductSearch)

	mux.HandleFunc("/user/mypage", routes.Mypage)

	server := &http.Server{
		Addr:  config.Address,
		Handler:mux,
	}

	server.ListenAndServe()
}
