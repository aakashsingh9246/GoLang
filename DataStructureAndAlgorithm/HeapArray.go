package main 

import(
	"fmt"
)

var heapArray []int
var track=0

func main(){
	insert(5)
	insert(9)
	insert(10)
	delete()
	delete()
	delete()
	fmt.Println(heapArray, track)
}

func insert(val int){
	heapArray = append(heapArray, val)
	track++
	if track-1==0{
		return
	}else{
		heapifyDtoU(track-1)
	}
}

func delete()int{
	temp:=heapArray[0]
	heapArray[0]=heapArray[track-1]
	heapArray = heapArray[:track-1]
	track--
	fmt.Printf("%v is deleted", temp)
	fmt.Println()
	heapifyUtoD(0)
	return temp
}

func heapifyUtoD(i int){
	l:=i*2+1
	r:=i*2+2
	if l<track && r<track{
		if heapArray[i]<heapArray[l] || heapArray[i]<heapArray[r]{
			if heapArray[l]>heapArray[r]{
				heapArray[i],heapArray[l]=heapArray[l],heapArray[i]
				heapifyUtoD(l)
			}else{
				heapArray[i],heapArray[r]=heapArray[r],heapArray[i]
				heapifyUtoD(r)
			}
		}
	}
}

func heapifyDtoU(i int){
	if i!=0{
		var pi int
	if i%2==0{
		pi=(i/2)-1
	}else{
		pi= i/2
	}
	if heapArray[pi]<heapArray[i]{
		heapArray[pi],heapArray[i] = heapArray[i],heapArray[pi]
	}
	heapifyDtoU(pi)
	}
}