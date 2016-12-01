// This files provides utilities for parsing the text files that contain the YouTube data
// it will use simple regex to split the values and then ultimately join them together into nodes
// then it will take all of the nodes and create a graph

package main

import (
	"os"
	"bufio"
	"strings"
	"errors"
	"strconv"
	"fmt"
)

// returns a new graph and the value of k (size of S)
func createGraphFromFile(fileName string) (graph, int) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)


	g := NewGraph()

	scanner.Scan()
	k := firstLine(scanner.Text())

	for scanner.Scan() {
		err, n := createNodeFromText(scanner.Text())
		if err != nil {
			continue
		}
		g.InsertNode(n)
	}

	return g, k
}

func createNodeFromText(line string) (error, node) {
	values := strings.Split(line, " ")

	if len(values) < 3 {
		e := errors.New("incomplete data")
		return e, node{}
	}

	id := values[0]
	toNode := values[1]
	probability, err := strconv.ParseFloat(values[2], 64)
	if err != nil {
		fmt.Println(err.Error())
	}


	edges := make([]edge, 0)

	e := edge{fromNode:id, toNode:toNode, probability:probability}
	edges = append(edges, e)

	n := node{id:id, edges:edges}

	return nil, n
}

// I really only need k (The size of S)
func firstLine(line string) int {
	values := strings.Split(line, " ")
	k, err := strconv.Atoi(values[2])
	if err != nil {
		fmt.Println(err.Error())
	}
	return k
}



