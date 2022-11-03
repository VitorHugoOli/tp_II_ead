package graphs

import (
	"II/analysis"
)

// Clean all edges from a graph
func (g *Graph) CleanEdges() {
	for i := 0; i < g.Nodes.Len; i++ {
		from := i
		for 0 < g.Nodes.At(from).Len {
			to := g.Nodes.At(from).At(0)
			g.RemoveEdge(from, to)
			g.RemoveEdge(to, from)
			analysis.Analysis.CurrentMeasure.Iterations++
		}
	}
}
