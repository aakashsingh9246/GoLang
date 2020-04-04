package main

import (
	"fmt"
)

// get max n put at end

func main() {
	data := []int{4,5,6,7,2,1,8,9}
	SelectionSort(data)
	fmt.Println(data)	
}

func SelectionSort(arr []int){
	size := len(arr)
	var max int
	for i:=0;i<size-1;i++{
		max=0
		for j:=1;j<size-i;j++{
			if arr[j]>arr[max]{
				max = j
			}
		}
		arr[size-i-1],arr[max] = arr[max],arr[size-i-1]
	}
}
