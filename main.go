package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var name interface{}
var subject interface{}
var message interface{}


func RenderTemplate(w http.ResponseWriter, url string){
	tmpl,_:=template.ParseFiles(url)
	tmpl.Execute(w,nil)
}
func Home(w http.ResponseWriter, r* http.Request){
	RenderTemplate(w,"forms.html")
	
	if r.Method=="POST"{
		name=r.FormValue("email")
		subject=r.FormValue("subject")
		message=r.FormValue("message")
		fmt.Println("Email is ",name)
		fmt.Println("Subject is ",subject)
		fmt.Println("Message is ",message)
	}
}
func main() {
	http.HandleFunc("/",Home)
	http.ListenAndServe(":8000",nil)

}