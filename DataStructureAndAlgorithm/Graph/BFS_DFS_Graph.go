package main

import (
	"fmt"
)

type edge struct {
	source byte
	dest   byte
	cost   int
	next   *edge
}

type edgeList struct {
	head *edge
}

type graph struct {
	count      int
	vertex     []byte
	vertexList map[byte]*edgeList
}

func (g *graph) Init(cnt int, vertex []byte) {
	g.count = cnt
	g.vertex = vertex
	g.vertexList = make(map[byte]*edgeList, cnt)
	for i := 0; i < cnt; i++ {
		g.vertexList[vertex[i]] = new(edgeList)
		g.vertexList[vertex[i]].head = nil
	}
}

func (g *graph) AddEdge(s, d byte, cost int) {
	edge := &edge{s, d, cost, g.vertexList[s].head}
	g.vertexList[s].head = edge
}

func (g *graph) AddBiDirEdge(s, d byte, cost int) {
	g.AddEdge(s, d, cost)
	g.AddEdge(d, s, cost)
}

func (g *graph) Print() {
	for i := 0; i < g.count; i++ {
		ad := g.vertexList[g.vertex[i]].head
		if ad != nil {
			fmt.Print("vertex ", string(ad.source), " ----> ")
			for ad != nil {
				fmt.Print("[", string(ad.dest), " ", ad.cost, "]")
				ad = ad.next
			}
			fmt.Println()
		}
	}
}

//DFS traversal of graph using stack

func (g *graph) DFSStack() {
	var curr byte
	visited := make(map[byte]bool)
	stk := new(Stack)
	visited[g.vertex[0]] = true
	stk.Push(g.vertex[0])
	for len(stk.stack) != 0 {
		curr = stk.Pop()
		fmt.Print(string(curr))

		head := g.vertexList[curr].head
		for head != nil {
			if !visited[head.dest] {
				visited[head.dest] = true
				stk.Push(head.dest)
			}
			head = head.next
		}
	}

}

type Stack struct {
	stack []byte
}

func (s *Stack) Push(c byte) {
	s.stack = append(s.stack, c)
}

func (s *Stack) Pop() byte {
	if len(s.stack) == 0 {
		fmt.Println("Stack is empty")
		return 0
	}
	val := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return val
}

func (s *Stack) isEmpty() bool {
	return len(s.stack) == 0
}

//DFS traversal of graph using recursion

func (g *graph) DFSRecursive() {
	visited := make(map[byte]bool)
	for i := 0; i < g.count; i++ {
		if !visited[g.vertex[i]] {
			visited[g.vertex[i]] = true
			g.dfsRcursiv(g.vertex[i], visited)
		}
	}
	fmt.Println()
}

func (g *graph) dfsRcursiv(ch byte, visit map[byte]bool) {
	curr := g.vertexList[ch].head
	fmt.Print(string(curr.source))
	for curr != nil {
		if !visit[curr.dest] {
			visit[curr.dest] = true
			g.dfsRcursiv(curr.dest, visit)
		}
		curr = curr.next
	}
}


// BFS Using queue

func (g *graph) BFS(){
	visited := make(map[byte]bool, g.count)
	for i:=0;i<g.count;i++{
		if !visited[g.vertex[i]]{
			g.bfsQueue(g.vertex[i], visited)
		}
	}
	
}

func (g *graph) bfsQueue(ch byte,visited map[byte]bool){
	var curr byte
	queue := new(Queue)
	queue.Enqueue(ch)
	visited[ch] = true
	for queue.size != 0{
		curr = queue.Dequeue()
		fmt.Print(string(curr))
		head := g.vertexList[curr].head
		for head != nil{
			if !visited[head.dest]{
				visited[head.dest] = true
				queue.Enqueue(head.dest)
			}
			head = head.next
		}
	}
	
}


// queue implementation 
const capacity int = 100

type Queue struct{
	size int
	data [capacity]byte
	start int
	end int
}

func (q *Queue) Enqueue(val byte){
	q.data[q.end] = val
	q.end = (q.end+1)%capacity
	q.size++ 
}

func (q *Queue) Dequeue ()byte{
	if q.size != 0{
		val := q.data[q.start]
		q.start = (q.start+1)%capacity
		q.size--
		return val
	}
	return 0
}


func main() {
	g := &graph{}
	ver := []byte{'A', 'B', 'C', 'D', 'E', 'F'}
	g.Init(6, ver)
	g.AddBiDirEdge('A', 'B', 1)
	g.AddBiDirEdge('A', 'C', 6)
	g.AddBiDirEdge('B', 'E', 1)
	g.AddBiDirEdge('B', 'D', 1)
	g.AddBiDirEdge('C', 'E', 5)
	g.AddBiDirEdge('D', 'E', 1)
	g.AddBiDirEdge('D', 'F', 4)
	g.AddBiDirEdge('E', 'F', 1)
	g.Print()
	g.DFSStack()
	fmt.Println()
	g.DFSRecursive()
	g.BFS()
}
