package main
import(
	"net/http"
	"log"
	"html/template"
)
var tpl *template.Template

func init(){
		tpl = template.Must(template.ParseGlob("templates/*"))
}

func main(){
	http.HandleFunc("/",index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, req *http.Request){
	err := tpl.ExecuteTemplate(w,"index.html",nil)
	if err!=nil{
		http.Error(w,err.Error(),500)
		log.Fatalln(err)
	}
}