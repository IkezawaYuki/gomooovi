package routes

import (
	"gomooovi/models"
	"net/http"
	"strconv"
)

func Mypage(w http.ResponseWriter, r *http.Request){

	ses, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		// user に紐づくデータも一緒に渡したい。Mapにすれば良いらしい。
		user, _ := models.UserByID(strconv.Itoa(ses.UserId))

		data := map[string]interface{}{"user":user, "dummy": "1"}

		generateHTML(w, data, "users/mypage")
	}

}

