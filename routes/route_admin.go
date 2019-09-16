package routes

import (
	"fmt"
	"gomooovi/models"
	"net/http"
)

func Admin(w http.ResponseWriter, r *http.Request){
	products, err_ := models.GetProductAll(20)
	if err_ != nil {
		fmt.Println(err_)
	}

	data := map[string]interface{}{"products":products}
	generateHTML(w, data, "admin/layout", "admin/product")

}