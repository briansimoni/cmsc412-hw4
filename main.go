package main

import (
	"fmt"
	"os"
	"strconv"
)


// the main function will create a graph object
// read the graph.txt file
func main() {

	g, k := createGraphFromFile("graph.txt")
	S, E := g.influenceMaximization(k)
	fmt.Println(S)
	fmt.Println(E)

	file, err := os.Create("im.txt")
	if err != nil {
		fmt.Println(err)
	}
	var s string
	for _, v := range S {
		s += v
	}
	s += "\n"
	s += string(strconv.FormatFloat(E, 'f', -4, 64))
	file.WriteString(s)

}