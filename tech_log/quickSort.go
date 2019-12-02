package main
import(
	"fmt"

)

func main(){
	a:=[]int{8,9,4,7,1,3,6,2}
	b:=len(a)-1
	quickSort(a,0,b)
	fmt.Println(a)
	a=append(a,5)
	getPartition(a,0,len(a)-1)
	fmt.Println(a)
}



func quickSort(arr []int, start int, end int){
	if start<end{
	pIndex := getPartition(arr, start, end)
	quickSort(arr,start, pIndex-1)
	quickSort(arr,pIndex+1,end)
	}
}

func getPartition(arr []int, start int, end int)int{
	pivot:=arr[end]
	pIndex := start
	for i:=start;i<end;i++{
		if(arr[i]<pivot){
			arr[i],arr[pIndex]=arr[pIndex],arr[i]
			pIndex++
		}
	}
	arr[pIndex],arr[end]=arr[end],arr[pIndex]
	return pIndex
}


/*func maximumToys(prices []int32, k int32) int32 {
    quickSort(prices, 0, len(prices)-1)
    var count int32=0
    var sum int32=0
    for _,v:=range prices{
        sum+=v
        if (sum <= k){
            count++
        }else{
            break
        }
    }
return count
}
func quickSort(arr []int32, start int, end int){
    if start<end{
    pIndex := getPartition(arr, start, end)
    quickSort(arr,start, pIndex-1)
    quickSort(arr,pIndex+1,end)
    }
}

func getPartition(arr []int32, start int, end int)int{
    pivot:=arr[end]
    pIndex := start
    for i:=start;i<end;i++{
        if(arr[i]<pivot){
            arr[i],arr[pIndex]=arr[pIndex],arr[i]
            pIndex++
        }
    }
    arr[pIndex],arr[end]=arr[end],arr[pIndex]
    return pIndex
}
*/