package main
import (
	"fmt"
)

func main(){
	fmt.Printf("Enter String:")
	var str string
	fmt.Scanln(&str)
	var mymap = make(map[string]int)
	for _,v:=range str{
		if _,ok:=mymap[string(v)];ok{
			mymap[string(v)]++
		}else{
			mymap[string(v)]=1
		}
	}
	result:=make([]string,len(str))
	permuteutil(mymap, result, 0)
}

func permuteutil(mymap map[string]int, result []string, level int){
	if level==len(result){
		fmt.Println(result)
		return
	}
	for k,v:=range mymap{
		if(v==0){
			continue
		}
		result[level]=k
		mymap[k]--
		permuteutil(mymap,result,level+1)
		mymap[k]++
	}
}