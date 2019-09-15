package routes

import (
	"fmt"
	"gomooovi/models"
	"log"
	"net/http"
	"strconv"
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
		fmt.Println("generate")
		generateHTML(w, product, "layouts/layout", "layouts/private.navbar", "reviews/new")
	}
}

func ReviewCreate(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}
	vals := r.URL.Query()
	productId := vals.Get("product")

	reviewComment := r.PostFormValue("review")
	rateStr := r.PostFormValue("rate")
	rate, err := strconv.Atoi(rateStr)

	ses, err := session(w, r)
	if err != nil{
		http.Redirect(w, r, "/login", 302)
	}else{
		user, err := models.UserByID(strconv.Itoa(ses.UserId))
		if err != nil{
			fmt.Println(err)
		}
		err = user.PostReview(productId, reviewComment, rate)

		vals.Set("id", productId)
		http.Redirect(w, r, "/products/show?id="+productId, 302)
	}

}

