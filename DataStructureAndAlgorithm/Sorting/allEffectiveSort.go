package main

import(
	"fmt"
)

func main(){
	arr := []int{4,6,5,7,8,2,1,3,9}
	heapSort(arr)
	fmt.Println(arr)

}


func mergeSort(arr []int)[]int{
	if len(arr)>1 {
		m:=len(arr)/2
		left:=arr[:m]
		right:=arr[m:]		
		return merge(mergeSort(left),mergeSort(right))
	}
	return arr
}

func merge(left, right []int)[]int{
	fmt.Println(left,"->",right)
	var temp []int 
	i,j :=0,0
	for i<len(left) && j<len(right){
		if left[i]<right[j]{
			temp = append(temp,left[i])
			i++
		}else{
			temp = append(temp,right[j])
			j++
		}
	}
	for i<len(left){
		temp = append(temp, left[i])
		i++
	}
	for j<len(right){
		temp = append(temp, right[j])
		j++
	}
	fmt.Println(temp)
	return temp
}

func quickSort(arr []int, s,e int){
	if s<e{
		i := partition(arr,s,e)
		quickSort(arr,s,i-1)
		quickSort(arr,i+1,e)
	}
}

func partition(arr []int, l,h int)int{
	pivot:= arr[h]
	j:=l-1
	for i:=l;i<=h;i++{
		if arr[i]<=pivot{
			j++
			arr[i],arr[j] = arr[j],arr[i]
		}
	}
	
	return j
}

func heapSort(arr []int){
	var heapArray []int
	for _,v:=range(arr){
		heapArray =insertInHeap(v, heapArray)
	}
	for i,_:=range(heapArray){
		heapArray, arr[i] = deleteFromHeap(heapArray)
	}
	fmt.Println(arr)
}


func insertInHeap(a int, heapArray []int)[]int{
	if len(heapArray)>1{
		heapArray = append(heapArray, a)
		heapArray = heapifyDtoU(heapArray, len(heapArray)-1)
	}else{
		heapArray = append(heapArray,a)
	}
	return heapArray
}

func heapifyDtoU(heapArray []int, i int)[]int{
	var	pi int
	if i!=0{
		if i%2==0{
			pi= i/2-1
		}else{
			pi=i/2
		}
		if heapArray[pi]>heapArray[i]{
			heapArray[pi],heapArray[i]=heapArray[i],heapArray[pi]
			heapifyDtoU(heapArray,pi)
		}
	}
	return heapArray
}


func deleteFromHeap(heapArray []int)([]int, int){
	i:=len(heapArray)-1
	val := heapArray[0]
	if len(heapArray)>1{
		heapArray[0],heapArray[i]=heapArray[i],heapArray[0]
	}
	heapArray = heapArray[:i]
	if len(heapArray)>1{
		heapArray = heapifyUtoD(heapArray, 0)
	}
	return heapArray,val	
}

func heapifyUtoD(heapArray []int, i int)[]int{
	l:=i*2+1
	r:=i*2+2
	if l<len(heapArray) && r<len(heapArray){
		if heapArray[l]<heapArray[i]||heapArray[r]<heapArray[i]{
			if heapArray[r]<heapArray[l]{
				heapArray[i],heapArray[r] = heapArray[r],heapArray[i]
				heapifyUtoD(heapArray,r)
			}else{
				heapArray[i],heapArray[l] = heapArray[l],heapArray[i]
				heapifyUtoD(heapArray,l)
			}
		}
	}
	return heapArray
}