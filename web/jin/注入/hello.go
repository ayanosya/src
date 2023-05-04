package main

import (
	"html/template"
	"net/http"
)

//用文件终端go build ./exe运行

func handle(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t := template.New("hello.html")
	t.Funcs(template.FuncMap{
		"func": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	_, _ = t.ParseFiles("hello.html")
	//渲染模板
	str1 := "<script>alert(1)</script>"
	str2 := "<script>alert(2)</script>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}
func main() {
	http.HandleFunc("/hello", handle)
	http.ListenAndServe(":9090", nil)
}
