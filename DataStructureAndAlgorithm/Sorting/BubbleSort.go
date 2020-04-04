package main

import (
	"fmt"
)
// Bubble Sort is slowest algorithm to sort can be used when input data is less

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

// improved bubble sort checks if swap is performed in current pass and decides list is sorted or not

func BubbleSort(data []int, comp more){
	size := len(data)
	flag := true  // to check swap is performed 
	for i:=0;i<size-1 && flag ;i++{
		flag = false
		for j:=0;j<size-i-1;j++{
			if comp(data[j],data[j+1]){
				data[j],data[j+1] = data[j+1],data[j]
				flag = true
			}
		}
	}
}
