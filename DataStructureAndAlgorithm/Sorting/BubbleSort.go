package main

import (
	"fmt"
)

type more func(int,int)bool 

func main() {
	data := []int{4,53,2,325,252,4524,25425,544}
	BubbleSort(data, compare)
	fmt.Println(data)	
}

func compare(i,j int)bool{
	return i>j
}

func BubbleSort(data []int, comp more){
	size := len(data)
	for i:=0;i<size-1;i++{
		for j:=0;j<size-i-1;j++{
			if comp(data[j],data[j+1]){
				data[j],data[j+1] = data[j+1],data[j]
			}
		}
	}
}
