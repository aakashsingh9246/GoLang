package main

import (
	"fmt"
	"time"
)

// Suggestions from golang-nuts
// http://play.golang.org/p/Ctg3_AQisl

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func helloworld(t time.Time) {
	fmt.Printf("%v: Hello, World!\n", t)
}

func main() {
	doEvery(1000*time.Millisecond, helloworld)
}


func startLog(){
	for{
		time.Sleep(2*time.Second)
		updateLog()
	}
}


file, err := os.OpenFile("test.txt", os.O_RDWR | os.O_APPEND, 0666)     
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    defer file.Close()
      
	for i:=0;i<10;i++{
		lineintext+=1
		str := "line no "+strconv.Itoa(lineintext)
		fmt.Println(str)
		fmt.Fprintln(file,str)
		if err!=nil{
			log.Fatalf("Failed to update ",err)
			}	
		}