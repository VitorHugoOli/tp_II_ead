package graphs

import (
	"II/lists/static"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type AdjaencyList interface {
	static.StaticList[*static.StaticList[int]]
}

type AjancencyInternal interface {
	static.StaticList[int]
}

// Graph Create a Adjacency List with the StaticList
type Graph struct {
	Nodes         *static.StaticList[*static.StaticList[int]]
	lenNeighbours int
}

// NewGraph Create a new Adjacency List
func NewGraph(size int, lenNeighbours int) *Graph {
	return &Graph{
		Nodes:         static.NewStaticList[*static.StaticList[int]](size),
		lenNeighbours: lenNeighbours,
	}
}

// AddNode Add a new node to the Adjacency List
func (g *Graph) AddNode(value int) {
	item := static.NewStaticList[int](g.lenNeighbours)
	g.Nodes.Add(item)
}

// AddEdge Add a new edge to the Adjacency List
func (g *Graph) AddEdge(from, to int) {
	g.Nodes.At(from).Add(to)
}

// Search Edge in the Adjacency List
func (g *Graph) Search(from, to int) bool {
	return g.Nodes.At(from).Search(to)
}

// RemoveEdge Edge in the Adjacency List
func (g *Graph) RemoveEdge(from, to int) {
	g.Nodes.At(from).Remove(to)
}

// Read Graph from a file
func (g *Graph) Read(filename string) {
	//read from ./instances/graph.txt
	// file in format: Number Line, nextNode, nextNode, nextNode, ...
	// 0, 1, 2, 3
	// 1, 0, 2, 3

	openFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer openFile.Close()

	// read the file
	scanner := bufio.NewScanner(openFile)
	lineCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		// split the line
		split := strings.Split(line, ",")
		// add the node
		g.AddNode(lineCount)
		// add the edges
		for i := 0; i < len(split); i++ {
			edge, _ := strconv.Atoi(split[i])
			g.AddEdge(lineCount, edge)
		}
		lineCount++
	}
}

// GetNumberOfEdges Get the number of edges in the graph
func (g *Graph) GetNumberOfEdges() int {
	edges := 0
	for i := 0; i < g.Nodes.Len; i++ {
		edges += g.Nodes.At(i).Len
	}
	return edges / 2
}

// Print Graph
func (g *Graph) Print() {
	for i := 0; i < g.Nodes.Len; i++ {
		print(i, ": ")
		g.Nodes.At(i).Print()
	}
}
