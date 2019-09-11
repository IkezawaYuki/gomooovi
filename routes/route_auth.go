package routes

import (
	"fmt"
	"gomooovi/models"
	"log"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("hello")
	generateHTML(writer, nil, "auth/layout", "layouts/public.navbar", "auth/login")
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth/layout", "layouts/public.navbar", "auth/signup")
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	user, err := models.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		// todo ログインできません的なメッセージを出す。
		//danger("connot find user")
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Fatalln(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)

	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != http.ErrNoCookie {
		session := models.Session{Uuid: cookie.Value}
		err = session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}
