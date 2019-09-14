package routes

import (
	"fmt"
	"gomooovi/models"
	"net/http"
)

func ReviewNew(w http.ResponseWriter, r *http.Request){
	vals := r.URL.Query()
	product_id := vals.Get("product")
	product, _ := models.GetProduct(product_id)
	fmt.Println(product)
	_, err := session(w, r)
	if err != nil{
		http.Redirect(w, r, "/login", 302)
	}else{
		generateHTML(w, product, "layouts/layout", "layouts/private.navbar", "reviews/new")
	}
}

func ReviewCreate(w http.ResponseWriter, r *http.Request){

}