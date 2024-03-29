package datastruct

import (
	"fmt"
	"testing"
)

type node string

func (s node) GetId() string {
	return string(s)
}

func TestGraph(t *testing.T) {
	graph := NewGraph[node]()
	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddNode("E")

	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("C", "D")

	result, _ := graph.TopologicalSort()
	fmt.Printf("%v\n", result)

	graph.RemoveEdge("A", "C")
	result, _ = graph.TopologicalSort()
	fmt.Printf("%v\n", result)

	graph.RemoveNode("D")
	result, _ = graph.TopologicalSort()
	fmt.Printf("%v\n", result)

}
