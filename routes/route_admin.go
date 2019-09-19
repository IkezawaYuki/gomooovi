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