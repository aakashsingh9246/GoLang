package main

import (
	"fmt"
)

type TrieNode struct{
	child map[string]*TrieNode
	isEnd bool 
}

var root TrieNode


func main(){
	root.child = make(map[string]*TrieNode)
	root.isEnd = false
	add("Akash", &root);
	fmt.Println(search("Akash"))
}

func newNode()TrieNode{
	a:=TrieNode{}
	a.child = make(map[string]*TrieNode)
	a.isEnd = false
	return a
}


func add(a string, node *TrieNode){
	current := root
	for i:=0;i<len(a);i++{
		ch := string(a[i])
		childNode := node.child[ch]
		if childNode ==nil{
			temp := newNode()
			current.child[ch] = &temp
			//fmt.Println(current.child)
		}
		fmt.Println(current)
		current = *current.child[ch]
	}
	current.isEnd = true
	fmt.Println(current)

}

func search(str string)bool{
	current := root
	for i:=0;i<len(str);i++{
		ch := string(str[i])
		//fmt.Println(current, i)
		childNode := current.child[ch]
		if childNode ==nil{
			fmt.Println("I am called")
			return false
		}
		
		current = *current.child[ch]
		//fmt.Println(current)
	}
	fmt.Println(current)
	return current.isEnd
}


func deleteN(current *TrieNode, str string, i int)bool{
	if i==len(str){
		if !current.isEnd{
			return false
		}
		current.isEnd=false
		return len(current.child)==0
	}
	ch := string(str[i])
	childNode := current.child[ch]
	if childNode==nil{
		return false
	}

	deleteCurrent := deleteN(childNode, str, i+1)
	if (deleteCurrent){
		delete(current.child, ch)
		return len(current.child) == 0
	}
	return false
}
