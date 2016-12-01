package main

import (
	"fmt"
	"math/rand"
)

// ... could be a better choice for negative infinity
const negativeInfinity = -987654321.012345
const T = 20000

// a graph is simply a collection of nodes and edges (the node struct contains edge data)
type graph struct {
	nodes map[string]node
}

func NewGraph() graph {
	nodes := make(map[string]node)
	return graph{nodes: nodes}
}

func (g *graph) PrintGraphInformation() {
	for id, _ := range g.nodes {
		fmt.Println("Node", id)
		for i := 0; i < len(g.nodes[id].edges); i++ {
			fmt.Println(g.nodes[id].edges[i].toNode)
		}

	}
}

// function to see if the node is already in the graph
func (g *graph) IsInGraph(id string) bool {
	if _, ok := g.nodes[id]; ok {
		return true
	}
	return false
}

func (g *graph) InsertNode(n node) {
	if g.IsInGraph(n.id) {
		ref := g.nodes[n.id]
		ref.AddEdge(n.edges[0])
		g.nodes[n.id] = ref
	} else {
		g.nodes[n.id] = n
	}
}

func (g *graph) visitedNodes() int {
	total := 0
	for _, node := range g.nodes {
		if node.visited {
			total++
		}
	}
	return total
}


// iterate through the graph and return
// a slice containing all of the node ID's
func (g *graph) getNodeSet() []string {
	set := make([]string, 0)
	for node, _ := range g.nodes {
		set = append(set, node)
	}
	return set
}


// return all of the elements in V that are not in S
// (assumes S is a subset of V)
func setDifference(V, S []string) []string {
	sPrime := make([]string, 0)
	found := false
	for _, v := range V {
		for _, s := range S {
			if v == s {
				found = true
				break;
			}
		}
		if !found {
			sPrime = append(sPrime, v)
		}
		found = false
	}
	return sPrime
}


func (g *graph) initialize() {
	for id, node := range g.nodes {
		node.visited = false
		g.nodes[id] = node
	}

}


// return the float64 expected number of activated nodes
// based on the start set S (S is a slice of node ID's)
func (g *graph) f(S []string) float64 {
	sum := 0
	for i := 1; i <= T; i++ {
		g.initialize()
		// create empty queue
		queue := make([]string, 0)
		for _, id := range S {
			queue = append(queue, id)
			node := g.nodes[id]
			node.visited = true
			g.nodes[id] = node
		}
		for len(queue) != 0 {
			// dequeue
			u := g.nodes[queue[0]]
			queue = queue[1:]
			for i := 0; i < len(u.edges); i ++ {

				neighbor := g.nodes[u.edges[i].toNode]
				probability := u.edges[i].probability
				if (!g.nodes[neighbor.id].visited) && (rand.Float64() <= probability) {

					neighbor.visited = true
					g.nodes[neighbor.id] = neighbor
					queue = append(queue, neighbor.id)

				}
			}
		}
		sum += g.visitedNodes()
	}
	return float64(sum) / T
}

// given the size of the seed set k return
// the set of nodes that results in the approximated
// maximum possible activated nodes in graph g
func (g *graph) influenceMaximization(k int) ([]string, float64) {
	S := make([]string, 0)
	V := g.getNodeSet()
	for i := 1; i <= k; i++ {
		best_v := ""
		best_delta := negativeInfinity
		iterateSet := setDifference(V, S)
		for _, v := range iterateSet {

			delta_v := g.f(append(S,v)) - g.f(S)
			if delta_v > best_delta {
				best_v = v
				best_delta = delta_v
			}

		}
		if best_v != "" {
			S = append(S, best_v)
		}
	}
	return S, g.f(S)
}
