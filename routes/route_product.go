package routes

import (
	"encoding/json"
	"fmt"
	"gomooovi/models"
	"io"
	"net/http"
	"strconv"
	"time"
)

func ProductSearch(w http.ResponseWriter, r *http.Request){

	_, err := session(w, r)
	rank, _ := models.GetRanking()

	data := map[string]interface{}{"rank": rank, "result": "dummy"}
	if err != nil{
		generateHTML(w, data, "layouts/layout", "layouts/public.navbar", "products/search")
	}else{
		generateHTML(w, data, "layouts/layout", "layouts/private.navbar", "products/search")
	}
}

func ProductShow(w http.ResponseWriter, r *http.Request){
	vals := r.URL.Query()
	id := vals.Get("id")
	product, _ := models.GetProduct(id)
	review, _ := models.GetReviewAll(id)
	rank, _ := models.GetRanking()

	data := map[string]interface{}{"product": product, "review": review, "rank": rank}
	_, err := session(w, r)
	if err != nil{
		generateHTML(w, data, "layouts/layout", "layouts/public.navbar", "products/show")
	}else{
		generateHTML(w, data, "layouts/layout", "layouts/private.navbar", "products/show")
	}
}

func ProductSearchApi(w http.ResponseWriter, r *http.Request){
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse json
	var jsonBody map[string]string
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	searchWord := jsonBody["word"]

	products, err := models.SearchProduct(searchWord)
	if err != nil {
		return
	}

	js, err := json.Marshal(products)
	fmt.Println(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(js)

	w.Header().Set("Content-type", "application/json")
	w.Write(js)
}

func SaveProduct(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	title := r.PostFormValue("title")
	image_url := r.PostFormValue("image_url")
	detail := r.PostFormValue("detail")
	director := r.PostFormValue("director")
	opendate := r.PostFormValue("opendate")

	product := &models.Product{
		Title:title,
		ImageUrl:image_url,
		CreatedAt:time.Now(),
		UpdatedAt:time.Now(),
		Director:director,
		Detail:detail,
		OpenDate:opendate,
	}

	err = product.RegisterProduct()
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/admin", 302)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse json
	var jsonBody map[string]string
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	productId := jsonBody["product_id"]

	fmt.Println(productId)

	http.Redirect(w, r, "/admin", 302)
}