package main
import(
	"math"
	"fmt"
)

func main(){
	fmt.Println(substrCount("abcbaba"))
}


func substrCount(s string) int64 {
    curr_char := ""
    strSeq := ""
    index:= -1
    var count int64 = 0
    var arr []int64

    for i:=0;i<len(s);i++{
        if i==0 || string(s[i])!=curr_char{
            index++
            curr_char = string(s[i])
            strSeq += curr_char
            arr = append(arr, int64(0)) 
        }
        arr[index]++
    }
    for _,v:=range arr{
        count += (v*(v+1))/2
    }
    for i,v:=range strSeq{
        if i>0 && i<len(strSeq)-1{
            if arr[i]==1 && string(strSeq[i-1])==string(strSeq[i+1]){
                count+= int64(math.Min(float64(arr[i-1]),float64(arr[i+1])))
            }
        }
    }

    return count
}
