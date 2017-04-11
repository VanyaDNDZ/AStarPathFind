package astart

import (
	"container/heap"
	"image/gif"
	"math"
)

type Point struct {
	X int
	Y int
}

type Node struct {
	came_from    *Node
	CurrentPoint Point
	Index        int
	Priority     int
	GScore       int
	NodeType     string
	Visited      bool
	InRoute      bool
}

func NewNode(point Point, nodeType string) *Node {
	node := &Node{}
	node.CurrentPoint = point
	node.GScore = 65000
	node.Priority = 65000
	node.came_from = nil
	node.Visited = false
	node.NodeType = nodeType
	node.InRoute = false

	return node
}

func (n Node) eq(other Node) bool {
	return n.CurrentPoint == other.CurrentPoint
}

func heuristic(a, b Point) int {
	return int(math.Abs(float64(a.X-b.X)) - math.Abs(float64(a.Y-b.Y)))
}

func Astar(graph *Graph2d, start, end Node, anim *gif.GIF) (*Node, bool) {
	var current *Node
	//c := make(chan Node)
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	heap.Push(openSet, &start)
	closeSet := &PriorityQueue{}
	heap.Init(closeSet)

	for openSet.Len() > 0 {
		current = heap.Pop(openSet).(*Node)
		current.Visited = true
		if current.eq(end) {
			return current, true
		}
		heap.Push(closeSet, current)

		for _, neighbor := range graph.GetNeighbors(current) {
			if closeSet.Has(*neighbor) {
				continue
			}
			tentative_gScore := current.GScore + heuristic(current.CurrentPoint, neighbor.CurrentPoint)
			if !openSet.Has(*neighbor) {
				heap.Push(openSet, neighbor)
			} else if tentative_gScore >= neighbor.GScore {
				continue
			}
			neighbor.came_from = current
			neighbor.GScore = tentative_gScore
			neighbor.Priority = neighbor.GScore + heuristic(neighbor.CurrentPoint, end.CurrentPoint)
		}
		AddFrame(anim, graph)
	}
	return &end, false
}

func PrintPath(graph *Graph2d, node *Node, anim *gif.GIF) {
	current := node
	for true {
		current.InRoute = true
		(*graph)[current.CurrentPoint.X][current.CurrentPoint.Y] = current
		AddFrame(anim, graph)
		if current.came_from == nil {
			break
		}
		current = current.came_from
	}
}
