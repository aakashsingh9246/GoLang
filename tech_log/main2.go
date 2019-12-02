package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "strings"
    "strconv"
)

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
    right := make(map[int64]int64)
    left := make(map[int64]int64)
    for _,v := range arr{
    	right[v]++
    }
    fmt.Println(right)
    var ans int64 = 0
    
    for _,val := range arr{
    	var c1,c2 int64 = 0,0
    	if (int64(val)%r == 0){
    		c1 = left[int64(val)/r]
    	}
    	right[int64(val)]--
    	c2 = right[int64(val)*r]
    	ans += c1*c2
    	left[int64(val)]++
    }
    fmt.Println(right)
    return ans
}




func main() {
    content, err := ioutil.ReadFile("tempText.txt")
    if err!=nil{
    	log.Fatal(err)
    }
    //fmt.Println(string(content))
    stringArray := strings.Split(string(content), " ")
    //fmt.Println(stringArray)
    var int64Array []int64
    for _,v := range stringArray{
    	tempVal, err := strconv.ParseInt(v, 10, 64)
    	if err!=nil{
    		log.Fatal(err)
    	}
    	int64Array = append(int64Array, tempVal)
    }
    //fmt.Println(int64Array)
   outPut := countTriplets(int64Array, 10)
   fmt.Println(outPut)
}

