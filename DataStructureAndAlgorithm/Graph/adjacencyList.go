package main 

import (
	"fmt"
)

type edge struct{
	source 	byte
	dest	byte
	cost 	int
	next	*edge
}

type edgeList struct{
	head	*edge
}

type graph struct{
	count 		int
	vertex   	[]byte
	vertexList 	map[byte]*edgeList
}

func (g *graph)Init(cnt int, vertex []byte){
	g.count = cnt
	g.vertex = vertex
	g.vertexList = make(map[byte]*edgeList, cnt)
	for i:=0;i<cnt;i++{
		g.vertexList[vertex[i]]=new(edgeList)
		g.vertexList[vertex[i]].head=nil
	}
}

func (g *graph)AddEdge(s,d byte, cost int){
	edge := &edge{s,d,cost,g.vertexList[s].head}
	g.vertexList[s].head = edge
}

func (g *graph)AddBiDirEdge(s,d byte,cost int){
	g.AddEdge(s,d,cost)
	g.AddEdge(d,s,cost)
}

func (g *graph)Print(){
	for i:=0;i<g.count;i++{
		ad := g.vertexList[g.vertex[i]].head
		if ad!=nil{
			fmt.Print("vertex ",string(g.vertex[i])," is connected to:")
			for ad!=nil{
				fmt.Print("[",string(ad.dest)," ",ad.cost,"]")
				ad=ad.next
			}
			fmt.Println()
		}
	}
}


func main(){
	g:=&graph{}
	ver := []byte{'A','B','C','D','E','F','G','H'}
	g.Init(8,ver)
	g.AddBiDirEdge('A','F',10)
	g.AddBiDirEdge('A','C',5)
	g.AddBiDirEdge('B','H',19)
	g.AddBiDirEdge('B','F',9)
	g.AddBiDirEdge('B','G',1)
	g.AddBiDirEdge('C','B',6)
	g.AddBiDirEdge('D','A',11)
	g.AddBiDirEdge('D','G',3)
	g.AddBiDirEdge('E','F',7)
	g.AddBiDirEdge('E','H',8)
	g.AddBiDirEdge('G','E',12)
	g.AddBiDirEdge('H','C',15)
	g.Print()
}
