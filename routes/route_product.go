package routes

import (
	"gomooovi/models"
	"net/http"
)

func ProductSearch(w http.ResponseWriter, r *http.Request){

	_, err := session(w, r)
	if err != nil{
		generateHTML(w, nil, "layouts/layout", "layouts/public.navbar", "products/search")
	}else{
		generateHTML(w, nil, "layouts/layout", "layouts/private.navbar", "products/search")
	}
}

func ProductShow(w http.ResponseWriter, r *http.Request){
	vals := r.URL.Query()
	id := vals.Get("id")
	product, _ := models.GetProduct(id)
	review, _ := models.GetReviewAll(id)

	data := map[string]interface{}{"product": product, "review": review, "dummy": "1"}
	_, err := session(w, r)
	if err != nil{
		generateHTML(w, data, "layouts/layout", "layouts/public.navbar", "products/show")
	}else{
		generateHTML(w, data, "layouts/layout", "layouts/private.navbar", "products/show")
	}
}