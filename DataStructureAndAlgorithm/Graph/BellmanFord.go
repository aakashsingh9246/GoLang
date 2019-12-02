package main

import (
	"fmt"
	"math"
)

type edge struct{
	src 	int
	dest 	int
	weight 	int
}

type graph struct{
	vertices	int
	edges		int
	edge		[]edge 	
}

func(g *graph)createGraph(v,e int){
	g.vertices = v
	g.edges = e
	g.edge = make([]edge, e)
}

func printAns(parent, dist []int){
	for k,_:=range parent{
		fmt.Println(k, dist[k], parent[k])
	}
}

func bellManFord(g *graph, src int){
	dist:=make([]int, g.vertices)
	parent:=make([]int, g.vertices)
	for i,_:=range dist{
		dist[i]=int(math.MaxInt32)
		parent[i]= 0
	}
	dist[src]=0
	v:=g.vertices
	e:=g.edges
	for i:=1;i<v;i++{
		for j:=0;j<e;j++{
			src := g.edge[j].src
			dest:=g.edge[j].dest
			weight := g.edge[j].weight
			if dist[src]!=int(math.MaxInt32) && dist[dest]>dist[src]+weight{
				dist[dest] = dist[src]+weight
				parent[dest]= src
			}
		}
	}

	for j:=0;j<e;j++{
			src := g.edge[j].src
			dest:=g.edge[j].dest
			weight := g.edge[j].weight
			if dist[src]!=int(math.MaxInt32) && dist[dest]>dist[src]+weight{
				fmt.Println("Negative cycle")
				return
			}
		}

	printAns(parent,dist)
}


func main(){
	V := 5; // Number of vertices in graph 
    E := 8; // Number of edges in graph 
    graph := &graph{}
    graph.createGraph(V, E); 
  
    // add edge 0-1 (or A-B in above figure) 
    graph.edge[0].src = 0; 
    graph.edge[0].dest = 1; 
    graph.edge[0].weight = -1; 
  
    // add edge 0-2 (or A-C in above figure) 
    graph.edge[1].src = 0; 
    graph.edge[1].dest = 2; 
    graph.edge[1].weight = 4; 
  
    // add edge 1-2 (or B-C in above figure) 
    graph.edge[2].src = 1; 
    graph.edge[2].dest = 2; 
    graph.edge[2].weight = 3; 
  
    // add edge 1-3 (or B-D in above figure) 
    graph.edge[3].src = 1; 
    graph.edge[3].dest = 3; 
    graph.edge[3].weight = 2; 
  
    // add edge 1-4 (or A-E in above figure) 
    graph.edge[4].src = 1; 
    graph.edge[4].dest = 4; 
    graph.edge[4].weight = 2; 
  
    // add edge 3-2 (or D-C in above figure) 
    graph.edge[5].src = 3; 
    graph.edge[5].dest = 2; 
    graph.edge[5].weight = 5; 
  
    // add edge 3-1 (or D-B in above figure) 
    graph.edge[6].src = 3; 
    graph.edge[6].dest = 1; 
    graph.edge[6].weight = 1; 
  
    // add edge 4-3 (or E-D in above figure) 
    graph.edge[7].src = 4; 
    graph.edge[7].dest = 3; 
    graph.edge[7].weight = -3; 
  
    bellManFord(graph, 0); 
}
