package main

import (
	"os"
	"log"
	"fmt"
	"strconv"
	"time"
	"net/http"
	"github.com/gorilla/websocket"
)

var line_no int
var lastUpdateLine int

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func main(){
	go startLog()
	http.HandleFunc("/", index)
	http.HandleFunc("/show", showLog)
	http.ListenAndServe(":8080", nil)
	
}	

func startLog(){
	for {
		time.Sleep(10*time.Second)
		updateLog()
	}
}

func updateLog(){
	file, err := os.OpenFile("test.txt", os.O_RDWR | os.O_APPEND, 0644)
     
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    defer file.Close()


    for i:=0;i<10;i++{
    	line_no+=1
    	str:= "you are at "+ strconv.Itoa(line_no)
    	_,err:= fmt.Fprintln(file, str)
    	if err!=nil {
    		log.Fatal(err)
    	}
    }
}


func index(w http.ResponseWriter, req *http.Request){
	fmt.Fprintln(w, "Hello Guest!")
}

func showLog(w http.ResponseWriter, req *http.Request){

	upgrader.CheckOrigin = func (req *http.Request)bool {return true}
	ws, err:= upgrader.Upgrade(w, req, nil)
	if err!=nil{
		log.Fatal(err)
	}

	reader(ws)
}

func reader(conn *websocket.Conn){
	for {
		messageType, p, err:= conn.ReadMessage()
		if err!=nil{
			log.Fatal(err)
			return
		}
		fmt.Println(string(p))
		err = conn.WriteMessage(messageType, p)
		if err!=nil{
			log.Fatal(err)
			return
		}

	}
}

func processLog(w http.ResponseWriter, req *http.Request){
	for {
		time.Sleep(1*time.Second)
		if checkLine() {
			fmt.Println("Hi there!")
			file, err := os.Open("test.txt")     
    		if err != nil {
        		log.Fatalf("failed opening file: %s", err)
    		}
    		defer file.Close()

    		buf:=make([]byte,620)
    		stat, err:= os.Stat("test.txt")
    		start := stat.Size()-620
   			_, err = file.ReadAt(buf, start)
    		if err!= nil{
    			log.Fatal(err)
    			}
    		fmt.Fprintln(w, string(buf))
			}
		}
}

func checkLine() bool{

	if (lastUpdateLine < line_no){	
		lastUpdateLine = line_no
		return true
	}
	return false
}
