package main

import (
	"fmt"
	"github.com/VanyaDNDZ/AStarPathFind/astar"
	"github.com/VanyaDNDZ/AStarPathFind/maze"
	"image/gif"
	"os"
)

func main() {

	matrix := maze.GenerateMaze(50)

	graph, start_node, end_node := astart.BuildGraph(matrix)
	f, err := os.OpenFile("rgb.gif", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	anim := &gif.GIF{}
	astart.AddFrame(anim, graph)

	if result, ok := astart.Astar(graph, *start_node, *end_node, anim); ok {
		astart.PrintPath(graph, result, anim)
	} else {
		fmt.Println("No result found")
	}
	astart.SaveGif(f, anim)

}
