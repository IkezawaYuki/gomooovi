package routes

import (
	"fmt"
	"gomooovi/models"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request){
	vals := r.URL.Query()
	page := vals.Get("page")
	fmt.Println(page)
	products, err := models.GetProductAll(20)
	if err != nil {
		fmt.Println(err)
	}

	data := map[string]interface{}{"products":products}
	generateHTML(w, data, "admin/layout", "admin/product")
}

func SaveProduct(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	title := r.PostFormValue("title")
	detail := r.PostFormValue("detail")
	director := r.PostFormValue("director")
	opendate := r.PostFormValue("opendate")

	fmt.Println(title)
	fmt.Println(detail)
	fmt.Println(director)
	fmt.Println(opendate)

	http.Redirect(w, r, "/admin", 302)
}