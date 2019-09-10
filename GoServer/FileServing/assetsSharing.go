/*package main

import(
	"io"
	"net/http"
)

func main(){
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./Assets"))))
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080",nil)
}

func dog(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/Assets/toby.jpg">`)
}*/


package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./Assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/Assets/toby.jpg">`)
}