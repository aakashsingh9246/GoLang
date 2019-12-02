package main
import (
	"fmt"
)

func main(){
	str:= "abcbaba"
	arr :=make([][]int, len(str))
	var ans []string
	for i := range arr{
		arr[i] = make([]int, len(str))
	}
	for i:=0;i<len(str);i++{
		arr[i][i] = 1
		ans = append(ans, str[i:i+1])
	}
	for i:=0;i<len(str)-1;i++{
		if str[i] == str[i+1]{
			arr[i][i+1]=1
			ans = append(ans, str[i:i+2])
		}
	} 

	for gap:=2;gap<len(str);gap++{
		for i:=0;i<len(str)-gap;i++{
			j:= gap+i
			fmt.Println(i,j)
			fmt.Println(arr[i+1][j-1])
			if str[i]==str[j] {
				arr[i][j]=1
				ans = append(ans, str[i:j+1])
			}
		}
	}

	fmt.Println(ans)
}