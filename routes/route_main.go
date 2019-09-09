package routes

import (
	"fmt"
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

	// todo セッション管理
	if true{
		generateHTML(w, nil, "layouts/layout", "layouts/public.navbar", "products/index")
	}else{
		generateHTML(w, nil, "layouts/layout", "layouts/private.navbar", "products/index")
	}
}

func ProductSearch(w http.ResponseWriter, r *http.Request){

	// todo セッション管理
	if true{
		generateHTML(w, nil, "layouts/layout", "layouts/public.navbar", "products/search")
	}else{
		generateHTML(w, nil, "layouts/layout", "layouts/private.navbar", "products/search")
	}

}