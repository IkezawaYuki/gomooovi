package routes

import (
	"errors"
	"fmt"
	"gomooovi/models"
	"html/template"
	"net/http"
)





// HTMLの生成
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string){

	var files []string
	for _, file := range filenames{
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	fmt.Println(files)
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func Index(w http.ResponseWriter, r *http.Request){

	_, err := session(w, r)
	if err != nil{
		generateHTML(w, nil, "layouts/layout", "layouts/public.navbar", "products/index")
	}else{
		generateHTML(w, nil, "layouts/layout", "layouts/private.navbar", "products/index")
	}
}

func ProductSearch(w http.ResponseWriter, r *http.Request){

	_, err := session(w, r)
	if err != nil{
		generateHTML(w, nil, "layouts/layout", "layouts/public.navbar", "products/search")
	}else{
		generateHTML(w, nil, "layouts/layout", "layouts/private.navbar", "products/search")
	}
}

func ProductShow(w http.ResponseWriter, r *http.Request){

	_, err := session(w, r)
	if err != nil{
		generateHTML(w, nil, "layouts/layout", "layouts/public.navbar", "products/show")
	}else{
		generateHTML(w, nil, "layouts/layout", "layouts/private.navbar", "products/show")
	}
}

func session(w http.ResponseWriter, r *http.Request) (ses models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		ses = models.Session{Uuid: cookie.Value}
		fmt.Println(ses)
		if ok, _ := ses.Check(); !ok {
			err = errors.New("invalid error")
		}
	}
	return
}