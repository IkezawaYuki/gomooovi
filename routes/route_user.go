package routes

import (
	"gomooovi/models"
	"log"
	"net/http"
)

func Mypage(w http.ResponseWriter, r *http.Request){
	vals := r.URL.Query()
	uuid := vals.Get("id")
	user, err := models.UserByUUID(uuid)
	if err != nil {
		log.Fatalln(err)
	} else {
		_, err := session(w, r)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
		} else {
			// user に紐づくデータも一緒に渡したい。Mapにすれば良いらしい。
			generateHTML(w, &user, "layouts/layout", "layouts/private.navbar", "users/show")
		}
	}
}

