package main
import (
	"fmt"
)

type Node struct{
	data int
	next *Node
}

type Info struct{
	Head *Node
	size int	
} 

func (i *Info)create(data int){
	if i.Head == nil {
		newNode := Node{data, nil}
		i.Head = &newNode
		i.size++
	}
}

func (i *Info)push(data int){
	if i.Head!= nil{
		newNode := Node{data,i.Head}
		i.Head= &newNode
		i.size++
	}
}

func(i *Info)pop()int{
	if i.Head!=nil{
		temp:=i.Head
		i.Head = i.Head.next
		i.size--
		return temp.data
	}
	return 0
}

func (i Info)peak()int{
	return i.Head.data
}

func (i Info)isEmpty() bool{
	if i.Head == nil{
		return true
	}

	return false
}

func (i *Info)deleteList(){
	i.Head.next = nil
}

func(i Info)printList(){
	if(i.Head!=nil){
		temp := i.Head
		for temp.next!=nil{
			fmt.Println(temp.data)
			temp= temp.next
			if(temp.next == nil){
				fmt.Println(temp.data)
			}
		}
	}
}


func main(){
	a:=&Info{nil, 0}
	b:=&Info{nil,0}
	a.create(10)
	b.create(1)
	b.push(2)
	b.push(3)
	b.push(4)
	fmt.Println(a.size)
	b.printList()
	fmt.Println(b.isEmpty())
	a.pop()
	fmt.Println(a.isEmpty())

}