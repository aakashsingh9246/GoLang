package main

import (
	"fmt"
)

type TrieNode struct{
	node_map map[string]*TrieNode
	isEnd bool 
}

var root TrieNode


func main(){
	root.node_map = make(map[string]*TrieNode)
	root.isEnd = false
	add("Akash", &root);
	add("ABCD", &root);
	add("Santosh", &root);
	add("AkashRohan", &root);
	fmt.Println(search("Akash"))
	fmt.Println(deleteN(&root,"Akash", 0))
	fmt.Println(search("Akash"))
	fmt.Println(printAll("", &root))
}

func newNode()TrieNode{
	a:=TrieNode{}
	a.node_map = make(map[string]*TrieNode)
	a.isEnd = false
	return a
}


func add(a string, node *TrieNode){
	
	for i:=0;i<len(a);i++{
		ch := string(a[i])
		nextNode := node.node_map[ch]
		if nextNode ==nil{
			temp := newNode()
			node.node_map[ch] = &temp
		}
		node = node.node_map[ch]
	}
	node.isEnd = true
}

func search(str string)bool{
	current := root
	for i:=0;i<len(str);i++{
		ch := string(str[i])
		nextNode := current.node_map[ch]
		if nextNode ==nil{
			return false
		}
		
		current = *current.node_map[ch]
	}
	return current.isEnd
}


func deleteN(current *TrieNode, str string, i int)bool{
	if i==len(str){
		if !current.isEnd{
			return false
		}
		current.isEnd=false
		return len(current.node_map)==0
	}
	ch := string(str[i])
	childNode := current.node_map[ch]
	if childNode==nil{
		return false
	}

	deleteCurrent := deleteN(childNode, str, i+1)
	if (deleteCurrent){
		delete(current.node_map, ch)
		return len(current.node_map) == 0
	}
	return false
}

func printAll (str string,node *TrieNode)[]string{
	current := node
	var word []string 
	if current.isEnd{
		word = append(word, str)
	}
	if len(current.node_map) != 0{
		for k,_ := range current.node_map{
		word = append(word, printAll(str+k,current.node_map[k])...)
		}
	}

	return word
}