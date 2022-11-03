package main

import (
	"II/analysis"
	"II/graphs"
	"fmt"
)

/*
Contiguous static list - staticList
Contiguous static list with a pointer to the first element - staticListSimple
Contiguous static list with a pointer to the first and last element - staticListDouble

Dynamic linked list with a pointer to the first element - linkedList
Dynamic linked list with a pointer to the first and last element - linkedListSimples
Dynamic linked list with a pointer to the first and last element and a pointer to the previous element - linkedListFull

*/

/*
 */
func main() {
	analysis.Analysis = analysis.FullAnalysis.Static

	instances := []int{100, 200, 300, 400, 500}

	for _, instance := range instances {
		graph := graphs.NewGraph(instance, instance)
		graph.Read("instances/graph_" + fmt.Sprint(instance) + ".txt")
		analysis.Analysis.NewMeasure(instance, graph.GetNumberOfEdges())
		graph.CleanEdges()
		analysis.Analysis.EndMeasure()
	}

	for _, e := range analysis.Analysis.Measurements {
		fmt.Println(e.Iterations)
	}

	analysis.FullAnalysis.Plot()

}
