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
	http.HandleFunc("/", createFile)
	http.HandleFunc("/read", readData)
	http.HandleFunc("/update", update)
	http.ListenAndServe(":8080", nil)
}

func createFile(w http.ResponseWriter, req *http.Request){
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


}

func update(w http.ResponseWriter, req *http.Request){
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
}

func readData(w http.ResponseWriter, req *http.Request){
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