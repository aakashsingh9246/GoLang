package main

import (
	"io"
	"net/http"
	"os"
	"log"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogFile)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!--image doesn't serve-->
	<img src="/toby.jpg">
	`)
}

func dogFile(w http.ResponseWriter, req *http.Request){
	f,err:=os.Open("toby.jpg")
	if err!=nil{
		log.Fatalln(err)
	}
	defer f.Close()
	io.Copy(w,f)
}
