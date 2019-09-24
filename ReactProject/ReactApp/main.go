package main

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
)

type Data struct{
	Val int `json: "val"`
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/start", createFile)
	http.HandleFunc("/update", update)
	http.HandleFunc("/read", readData)
	http.HandleFunc("/getFile", showFile)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}


func index(w http.ResponseWriter, req *http.Request){
	fmt.Fprintln(w,"Welcome to index page!")
}


func createFile(w http.ResponseWriter, req *http.Request){
	setupResponse(&w, req)
	jsonFile, err := os.Create("myData.json")
	if err!=nil{
		fmt.Println("Error in creating JSON file: ", err)
		return
	}
	defer jsonFile.Close()

	post:= Data{1}

	updatedFile, err:= json.Marshal(&post)
   	if err != nil{
   		fmt.Println("Error in marshalling: ", err)
   		return
   	}

   	err = ioutil.WriteFile("myData.json", updatedFile, 0644)
   	if err != nil{
   		fmt.Println("Error in writing json file: ", err)
   		return
   	}
   	fmt.Fprintln(w,post.Val)

}

func update(w http.ResponseWriter, req *http.Request){
	setupResponse(&w, req)
	jsonFile, err := os.Open("myData.json")
	if err!=nil{
		fmt.Println("Error in opening JSON file: ", err)
		return
	}
 	defer jsonFile.Close()

 	jsonData, err:= ioutil.ReadAll(jsonFile)
 	if err!=nil{
		fmt.Println("Error in reading JSON file: ", err)
		return
	}
    var post Data
   	json.Unmarshal(jsonData, &post)
   	
   	post.Val++
   	updatedFile, err:= json.Marshal(&post)
   	if err != nil{
   		fmt.Println("Error in marshalling: ", err)
   		return
   	}

   	err = ioutil.WriteFile("myData.json", updatedFile, 0644)
   	if err != nil{
   		fmt.Println("Error in writing json file: ", err)
   		return
   	}
   	fmt.Fprintln(w,post.Val)
}

func readData(w http.ResponseWriter, req *http.Request){
	setupResponse(&w, req)
	jsonFile, err := os.Open("myData.json")
	if err!=nil{
		fmt.Println("Error in opening JSON file: ", err)
		return 
	}
 	defer jsonFile.Close()

 	jsonData, err:= ioutil.ReadAll(jsonFile)
 	if err!=nil{
		fmt.Println("Error in reading JSON file: ", err)
		return 
	}
    var post Data
   	json.Unmarshal(jsonData, &post)
   	fmt.Fprintln(w,"Data ", post.Val)
}

func showFile(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	update(w,req)
	http.ServeFile(w, req, "myData.json")
}