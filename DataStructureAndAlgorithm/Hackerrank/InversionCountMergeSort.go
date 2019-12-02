package main
import (
	"fmt"
)

func main(){
	myCount := countInversions([]int{7, 5, 3, 1})
	fmt.Println(myCount)
}


func countInversions(arr []int)int{
	temp := make([]int, len(arr))
	count := mergeSort(arr, temp,0, len(arr)-1)
	return count
}

func mergeSort(arr,temp []int, l,r int) int{

	countInv := 0
	if (l<r){
		mid := (l+r)/2
		countInv = mergeSort(arr, temp,l, mid)
		countInv+= mergeSort(arr, temp, mid+1, r)
		countInv+= merge(arr, temp, l, mid+1, r)
	}
	return countInv
}


func merge(a,t []int, l,m,r int)int{
	i:=l
	j:=m
	k:=l
	count:=0
	for (i<=m-1 && j<=r){
		if a[i]<=a[j]{
			t[k]=a[i]
			i++
			k++
		}else{
			t[k] = a[j]
			j++
			k++
			count += (m-i)
		}
	}
	for i<=m-1{
		t[k]=a[i]
		i++
		k++
	}
	for j<=r{
		t[k]=a[j]
		j++
		k++
	}
	for i:=l;i<=r;i++{
		a[i]=t[i]
	}
	return count
}




