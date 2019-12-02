package main 

import (
	"fmt"
)


type Node struct{
	data int
	left *Node
	right *Node
}

type Tree struct{
	root *Node
}

func (t *Tree)AddNode(val int){
	t.root = AddNodeUtil(t.root, val)
}

func AddNodeUtil(n *Node, val int)*Node{
	if n==nil{
		new := &Node{}
		new.data = val
		return new
	}
	if val<n.data{
		n.left = AddNodeUtil(n.left, val)
	} else {
		n.right = AddNodeUtil(n.right, val)
	}
	return n
}

func (t *Tree)Find(val int)bool{
	curr := t.root
	for curr != nil{
		if curr.data == val{
		return true
		}
		if val > curr.data{
			curr = curr.right
		}else{
			curr = curr.left
		}
	}
	return false
}

func (t *Tree)FindMin()(int,bool){
	curr := t.root
	if curr==nil{
		return 0,false
	}
	for curr.left !=nil{
		curr = curr.left
	}
	return curr.data,true
}

func main(){
	t := new(Tree)
	arr := []int{5,4,3,7,8,9,1,3,34,23,14,67,89,90}
	for _,v:=range arr{
		t.AddNode(v)
	}

	fmt.Println(t.FindMin())
}


