package routes

import (
	"fmt"
	"gomooovi/models"
	"log"
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


func AdminLogin(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth/layout", "layouts/public.navbar", "admin/login")
}

func AdminUsers(writer http.ResponseWriter, request *http.Request) {
	users, err := models.Users()
	if err != nil {
		fmt.Println(err)
	}
	admin, err := models.AdminUsers()
	if err != nil {
		fmt.Println(err)
	}
	data := map[string]interface{}{"users":users, "admin":admin}

	fmt.Println(users)
	fmt.Println(admin)
	generateHTML(writer, data, "admin/layout", "layouts/public.navbar", "admin/users")
}

func AdminAuthenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	user, err := models.AdminUserByEmail(r.PostFormValue("email"))
	if err != nil {
		// todo ログインできません的なメッセージを出す。
		//danger("connot find user")
	}
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateAdminSession()
		if err != nil {
			log.Fatalln(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie_admin",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/admin", 302)
	} else {
		http.Redirect(w, r, "/admin/login", 302)
	}
}

func AdminSignupAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	user := models.AdminUser{
		Nickname: r.PostFormValue("nickname"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err = user.Create(); err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/admin/login", 302)
}



func AdminLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie_admin")
	if err != http.ErrNoCookie {
		session := models.AdminSession{Uuid: cookie.Value}
		err = session.DeleteByUUID()
	}
	http.Redirect(w, r, "/", 302)
}

