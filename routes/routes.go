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
		files = append(files, fmt.Sprintf("views/layouts/%s.html", file))
	}
	fmt.Println(files)
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}


func Index(w http.ResponseWriter, r *http.Request){
	fmt.Println("aa")
	generateHTML(w, nil, "layout")
}