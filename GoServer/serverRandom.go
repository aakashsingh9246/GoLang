
package main
 
import (
    "io/ioutil"
    "log"
    "fmt"
    "os"
)
 
func Beginning() {
    // Read Write Mode
    file, err := os.OpenFile("test.txt", os.O_RDWR, 0644)
     
    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }
    defer file.Close()
     

for i:=0;i<10;i++{

    fmt.Fprintln(file, "Log statement ",i)
}

}
 
func ReadFile() {
    data, err := ioutil.ReadFile("test.txt")
    if err != nil {
        log.Panicf("failed reading data from file: %s", err)
    }
    fmt.Printf("\nFile Content: %s", data)  
}
 
func main() {
    fmt.Printf("\n######## Read file #########\n")
    ReadFile()
     
    fmt.Printf("\n\n######## WriteAt #########\n")
    Beginning()
     
    fmt.Printf("\n\n######## Read file #########\n")
    ReadFile()
}