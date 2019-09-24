package routes

import (
	"fmt"
	"gomooovi/models"
	"net/http"
	"strconv"
)

func Admin(w http.ResponseWriter, r *http.Request){
	vals := r.URL.Query()
	page := vals.Get("page")
	if page == ""{
		page = "1"
	}
	start, _ := strconv.Atoi(page)
	products, err := models.GetProductAll(start * 20-20)
	if err != nil {
		fmt.Println(err)
	}

	data := map[string]interface{}{"products":products}
	generateHTML(w, data, "admin/layout", "admin/product")
}

