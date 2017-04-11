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
	node.GScore = 0
	node.Priority = 0
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
	difX := math.Abs(float64(a.X-b.X))
	difY := math.Abs(float64(a.Y-b.Y))
	return - int( difY + difX )
}

func gscore(node *Node) int {
	score := 0
	current := node
	for true {
		if current.came_from == nil {
			break
		}
		score += heuristic(current.CurrentPoint, current.came_from.CurrentPoint)
		current = current.came_from

	}
	return score
}

func Astar(graph *Graph2d, start, end Node, anim *gif.GIF) (*Node, bool) {
	var current *Node
	//c := make(chan Node)
	openSet := &PriorityQueue{}
	closeSet := &PriorityQueue{}
	heap.Init(openSet)
	heap.Init(closeSet)

	heap.Push(openSet, &start)


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
			tentative_gScore := gscore(neighbor)
			if !openSet.Has(*neighbor) {
				neighbor.came_from = current
				neighbor.GScore = tentative_gScore
				neighbor.Priority = neighbor.GScore + heuristic(neighbor.CurrentPoint, end.CurrentPoint)
				heap.Push(openSet, neighbor)
			} else if tentative_gScore >= neighbor.GScore {
				continue
			}
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
