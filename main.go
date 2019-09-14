package main

import (
	"gomooovi/routes"
	"net/http"
	"time"

	_ "gomooovi/models"
)

func main(){

	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", routes.Index)

	mux.HandleFunc("/login", routes.Login)
	mux.HandleFunc("/signup", routes.Signup)
	mux.HandleFunc("/authenticate", routes.Authenticate)
	mux.HandleFunc("/logout", routes.Logout)
	mux.HandleFunc("/signup_account", routes.SignupAccount)

	//mux.Handle("/login", &routes.TemplateHandler{Filenames: []string{"auth/layout", "layouts/public.navbar", "auth/login"}})


	mux.HandleFunc("/products/search", routes.ProductSearch)
	mux.HandleFunc("/products/show", routes.ProductShow)

	mux.HandleFunc("/reviews/new", routes.ReviewNew)
	mux.HandleFunc("/reviews/create", routes.ReviewCreate)

	mux.HandleFunc("/users/mypage", routes.Mypage)

	server := &http.Server{
		Addr:  config.Address,
		Handler:mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
