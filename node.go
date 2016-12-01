package main

import (
	"reflect"
)

type node struct {
	id       string
	edges    []edge
	visited bool

}

func (n *node) AddEdge(e edge) {
	for i := range n.edges {
		if reflect.DeepEqual(e, n.edges[i]) {
			return
		}
	}

	n.edges = append(n.edges, e)
}
