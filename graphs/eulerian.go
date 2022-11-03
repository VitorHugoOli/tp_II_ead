package graphs

import "II/analysis"

// IsEulerianCircuit Check if the graph is Eulerian Circuit
func (g *Graph) IsEulerianCircuit() bool {
	if !g.IsConnected() {
		return false
	} else {
		for i := 0; i < g.Nodes.Len; i++ {
			if g.Nodes.At(i).Len%2 != 0 {
				return false
			}
		}
		return true
	}
}

// IsBridge Check if is a bridge
func (g *Graph) IsBridge(from, to int, ignoreNodes ...[]int) bool {
	g.RemoveEdge(from, to)
	g.RemoveEdge(to, from)
	isBridge := !g.IsConnected(ignoreNodes[0])
	g.AddEdge(from, to)
	g.AddEdge(to, from)
	return isBridge
}

// IsConnected Check if the graph is connected
func (g *Graph) IsConnected(ignoreNodes ...[]int) bool {
	visited := make([]bool, g.Nodes.Len)
	g.DFS(0, visited)

	analysis.Analysis.CurrentMeasure.Iterations++
	if ignoreNodes != nil {
		for _, e := range ignoreNodes[0] {
			analysis.Analysis.CurrentMeasure.Iterations++
			visited[e] = true
		}
	}

	for i := 0; i < len(visited); i++ {
		analysis.Analysis.CurrentMeasure.Iterations++
		if !visited[i] {
			return false
		}
	}
	return true
}

// DFS Depth First Search
func (g *Graph) DFS(node int, visited []bool) {
	visited[node] = true
	for i := 0; i < g.Nodes.At(node).Len; i++ {
		analysis.Analysis.CurrentMeasure.Iterations++
		if !visited[g.Nodes.At(node).At(i)] {
			g.DFS(g.Nodes.At(node).At(i), visited)
		}
	}
}

// printPath Print the path
func printPath(path []int) {
	for i := 0; i < len(path); i++ {
		analysis.Analysis.CurrentMeasure.Iterations++
		print(path[i])
		if i != len(path)-1 {
			print(" -> ")
		}
	}
	println()
}

// Copy the graph
func (g *Graph) Copy() *Graph {
	graph := NewGraph(g.Nodes.Len, g.lenNeighbours)
	for i := 0; i < g.Nodes.Len; i++ {
		graph.AddNode(i)
		for j := 0; j < g.Nodes.At(i).Len; j++ {
			analysis.Analysis.CurrentMeasure.Iterations++
			graph.AddEdge(i, g.Nodes.At(i).At(j))
		}
	}
	return graph
}

// IsValidEdge Check if the edge is valid
func (g *Graph) IsValidEdge(from, to int, ignoreNodes ...[]int) bool {
	analysis.Analysis.CurrentMeasure.Iterations++
	if g.Nodes.At(from).Len == 1 {
		return true
	} else {
		isBridge := g.IsBridge(from, to, ignoreNodes[0])
		return !isBridge
	}
}

// EulerianCircuit Find the Eulerian Circuit and return the path
func (g *Graph) EulerianCircuit() []int {
	var path []int
	graph := g.Copy()
	graph.EulerianCircuitRec(0, &path, nil)
	//printPath(path)
	return path
}

// EulerianCircuitRec Find the Eulerian Circuit and return the path
func (g *Graph) EulerianCircuitRec(node int, path *[]int, excludeNodes *[]int) {
	analysis.Analysis.CurrentMeasure.Iterations++
	if excludeNodes == nil {
		excludeNodes = &[]int{}
	}

	for i := 0; i < g.Nodes.At(node).Len; i++ {
		no := g.Nodes.At(node)
		to := no.At(i)
		analysis.Analysis.CurrentMeasure.Iterations++
		if g.IsValidEdge(node, to, *excludeNodes) {
			g.RemoveEdge(node, to)
			g.RemoveEdge(to, node)
			analysis.Analysis.CurrentMeasure.Iterations++
			if no.Len == 0 {
				*excludeNodes = append(*excludeNodes, node)
			}
			g.EulerianCircuitRec(to, path, excludeNodes)
		}
	}

	*path = append(*path, node)
}
