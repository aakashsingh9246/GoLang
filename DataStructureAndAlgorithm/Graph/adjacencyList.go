package main 

import (
	"fmt"
)

type edge struct{
	source 	int
	dest	int
	cost 	int
	next	*edge
}

type edgeList struct{
	head	map[byte]*edge
}

type graph struct{
	count 		int
	vertexList 	[]*edgeList
}

func (g *graph)Init(cnt int){
	g.count = cnt
	g.vertexList = make([]*edgeList, cnt)
	for i:=0;i<cnt;i++{
		g.vertexList[i]=new(edgeList)
		g.vertexList[i].head=nil
	}
}

func (g *graph)AddEdge(s,d byte, cost int){
	edge := &edge{s,d,c,g.vertexList[s].head}
	g.vertexList[s].head = edge
}

func (g *graph)Print(){
	for i:=0;i<g.count;i++{
		ad := g.vertexList[i].head
		if ad!=nil{
			fmt.Print("vertex",i,"is connected to:")
			for ad!=nil{
				fmt.Print("[",ad.dest,ad.cost,"]")
				ad=ad.next
			}
			fmt.Println()
		}
	}
}


func main(){
	
}

