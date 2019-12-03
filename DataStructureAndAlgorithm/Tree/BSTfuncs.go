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

//to add new node to tree

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

//to check whether given numbet is in tree

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

// to get min max

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

func (t *Tree)FindMax()(int,bool){
	curr := t.root
	if curr==nil{
		return 0,false
	}
	for curr.right !=nil{
		curr = curr.right
	}
	return curr.data,true
}

//to confirm whether tree is BST  

func (t *Tree)IsBST()bool{
	var c int
	return isBstNode2(t.root, &c)//isBstNode(t.root) // 
}

func isBstNode(n *Node)bool{
	if n.left!=nil{
		if n.data<n.left.data{
		return false
		}else{
			return isBstNode(n.left)
		}
	}
	if n.right!=nil {
		if n.data>n.right.data{
		return false
		}else{
			return isBstNode(n.right)
		}
	}
	return true
}

func isBstNode2(n *Node, count *int)bool{
	var ret bool
	if n!=nil{
	ret = isBstNode2(n.left, count)
	if !ret{
		return false
	}
	if *count>n.data{
		return false
	}
	*count = n.data
	ret = isBstNode2(n.right, count)
	if !ret{
		return false
	}
}
return true
}

//to delete given int from tree

func (t *Tree)deleteNode(val int){
	t.root = deleteNode(t.root, val)
}

func deleteNode(n *Node, val int)*Node{
	var temp = new(Node)
	if n==nil{
		return nil
	}
	if n.data==val{
		if n.left==nil && n.right==nil{
			return nil
		}
		if n.left==nil{
			temp = n.right
			return temp
		}
		if n.right==nil{
			temp = n.left
			return temp
		}
		min := getMin(n.right)
		n.data = min
		n.right = deleteNode(n.right,min)

	}else{
		if n.data<val{
			n.right = deleteNode(n.right,val)
		}else{
			n.left = deleteNode(n.left,val)
		}
	}
	return n
}


func getMin(curr *Node)int{
	if curr==nil{
		return 0
	}
	for curr.left !=nil{
		curr = curr.left
	}
	return curr.data
}

//least common ancestor

func (t *Tree)getLCA(a,b int)int{
	return getLCA(t.root, a,b)
}

func getLCA(n *Node, a,b int)int{
	if n==nil{
		return 0
	}
	if a<n.data && b<n.data{
		return getLCA(n.left,a,b)
	}
	if a>n.data && b>n.data{
		return getLCA(n.right,a,b)
	}
	return n.data
}

// delete node which comes out range

func (t *Tree)trimTree(a,b int){
	t.root = trimTree(t.root,a,b)
}

func trimTree(n *Node, min,max int)*Node{
	if n==nil{
		return nil
	}
	n.left = trimTree(n.left,min,max)
	n.right = trimTree(n.right, min,max)
	if n.data<min{
		return n.right
	}
	if n.data>max{
		return n.left
	}
	return n
}

// to print in range 

func (t *Tree)printInrange(a,b int){
	printInrange(t.root,a,b)
}

func printInrange(n *Node, min,max int){
	if n==nil{
		return
	}
	printInrange(n.left,min,max)
	if n.data>min && n.data<max{
		fmt.Print(n.data, " ")
	}
	printInrange(n.right,min,max)
}

//inorder traversal

func (t *Tree) inOrderTraversal(){
	inOrderTraversalNode(t.root)
}

func inOrderTraversalNode(n *Node){
	if n==nil{
		return
	}
	inOrderTraversalNode(n.left)
	fmt.Print(n.data," ")
	inOrderTraversalNode(n.right)
}

func main(){
	t := new(Tree)
	arr := []int{20,10,30,5,15,25,35,2,7,12,17,22,27,32,37}
	for _,v:=range arr{
		t.AddNode(v)
	}
	//t.deleteNode(91)
	//fmt.Println(t.FindMin())
	t.inOrderTraversal()
	fmt.Println()
	t.printInrange(6,26)
	fmt.Println()
	t.inOrderTraversal()
	fmt.Println()

	//fmt.Println(t.getLCA(9,1))
	//fmt.Println(t.IsBST())
}
