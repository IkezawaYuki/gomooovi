package routes

import (
	"errors"
	"fmt"
	"gomooovi/models"
	"html/template"
	"net/http"
	"strconv"
)




// HTMLの生成
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string){

	var files []string
	for _, file := range filenames{
		files = append(files, fmt.Sprintf("views/%s.html", file))
	}
	fm := template.FuncMap{
		"add":   models.Add,
		"rate": models.ReviewAverage,
	}
	t := template.New("content").Funcs(fm)
	t = template.Must(t.ParseFiles(files...))
	t.ExecuteTemplate(w, "layout", data)
}




func Index(w http.ResponseWriter, r *http.Request){

	vals := r.URL.Query()
	page := vals.Get("page")
	if page == ""{
		page = "1"
	}
	start, _ := strconv.Atoi(page)
	products, err_ := models.GetProductAll(start * 20-20)
	if err_ != nil {
		fmt.Println(err_)
	}

	rank, err_ := models.GetRanking()
	if err_ != nil {
		fmt.Println(err_)
	}

	data := map[string]interface{}{"products":products, "rank": rank}
	_, err := session(w, r)
	if err != nil{
		generateHTML(w, data, "layouts/layout", "layouts/public.navbar", "products/index")
	}else{
		generateHTML(w, data, "layouts/layout", "layouts/private.navbar", "products/index")
	}
}



func session(w http.ResponseWriter, r *http.Request) (ses models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		ses = models.Session{Uuid: cookie.Value}
		if ok, _ := ses.Check(); !ok {
			err = errors.New("invalid error")
		}
	}
	return
}