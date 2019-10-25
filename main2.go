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
    mymap := getMap(arr)
    var result []int64
    if r>1{
        for k,v:= range mymap{
            val1,ok:=mymap[k*r]
            val2,dk :=mymap[k*r*r]
            if ok && dk{
            	fmt.Println(k)
            	fmt.Println(v,val1,val2)
                result = append(result, v*val1*val2)
            }
        }
    } else if r==1 {
        for _,v:=range mymap{
            if v>=3{
                var val int64 = ((v)*(v-1)*(v-2))/6
                //fmt.Println(val)
                result = append(result, val)  
            }
        }
    }
    var outPut int64 = 0
    for _,ans := range result{
        outPut += int64(ans) 
    }
    fmt.Println(result)
    return outPut
}

func getMap(arr []int64) map[int64]int64{
     m := make(map[int64]int64)
    for _,v := range arr{
        if _,ok := m[int64(v)]; ok{
            m[int64(v)] = m[int64(v)]+int64(1)
        }else{
            m[int64(v)]=1
        }
    }
    fmt.Println(m)
 return m   
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
   outPut := countTriplets(int64Array, 3)
   fmt.Println(outPut)
}

