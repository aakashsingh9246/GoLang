package main

import(
	"fmt"
)


func main(){
	sortedArray:= sort([]int{5,4,7,6,2,3,1,9,8})
	fmt.Println(sortedArray)
}

func sort(arr []int)[]int{
	for i:=1;i<len(arr);i++{
		j:=i
		temp:=arr[i]
		for j>0 && arr[j-1]>temp{
			arr[j]=arr[j-1]
			j--
		}
		arr[j] = temp
	}
	return arr
}
