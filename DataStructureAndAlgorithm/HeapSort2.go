package main 

import(
	"fmt"
)

type Heap struct{
	size int
	arr []int
	isMin bool
}

func NewHeap(arrInput []int, isMin bool)*Heap{
	size := len(arrInput)
	arr:=[]int{1}
	arr = append(arr, arrInput...)
	h:=&Heap{size,arr,isMin}
	for i:=size/2;i>0;i--{
		h.proclateDown(i)
	}
	return h
}

func NewHeap2(isMin bool)*Heap{
	arr:=[]int{1}
	return &Heap{0,arr,isMin}
}

func (h *Heap)proclateDown(parent int){
	lChild := parent*2
	rChild := lChild+1
	small:=-1
	if lChild<=h.size{
		small=lChild
	}
	if rChild<=h.size && h.comp(lChild,rChild){
		small = rChild
	}
	if small!=-1 && h.comp(parent,small){
		h.swap(parent,small)
		h.proclateDown(small)
	}
}

func (h *Heap)proclateUp(child int){
	parent:= child/2
	if parent==0{
		return 
	}
	if h.comp(parent,child){
		h.swap(parent,child)
		h.proclateUp(parent)
	}
}

func (h *Heap)comp(i,j int)bool{
	if h.isMin{
		return h.arr[i]>h.arr[j]
	}
	return h.arr[i]<h.arr[j]
}

func (h *Heap)swap(i,j int){
	h.arr[i],h.arr[j] = h.arr[j],h.arr[i]
}

func (h *Heap)isEmpty()bool{
	return h.size == 0
}

func (h *Heap)peek()(int,bool){
	if h.isEmpty(){
		return 0,false
	}
	return h.arr[1],true
}

func (h *Heap)Size()int{
	return h.size
}

func (h *Heap)Add(val int){
	h.size++
	h.arr=append(h.arr,val)
	h.proclateUp(h.size)
}

func (h *Heap)Remove()(int,bool){
	if h.size==0{
		return 0,false
	}
	val:=h.arr[1]
	h.arr[1]=h.arr[h.size]
	h.arr=h.arr[:h.size]
	h.size--	
	h.proclateDown(1)
	return val,true
}

func HeapSort(inputArr []int){
	hp:=NewHeap(inputArr,true)
	for i,_:=range inputArr{
		inputArr[i],_=hp.Remove()
	}
}


func main(){
	a:=[]int{1,9,6,7,8,-1,-5,10,2}
	HeapSort(a)
	fmt.Println(a)
}