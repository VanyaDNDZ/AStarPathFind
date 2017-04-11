package astart

type Graph2d [][]*Node

func (graph *Graph2d) GetNeighbors(node *Node) []*Node {
	nodes := []*Node{}
	for row := node.CurrentPoint.X - 1; row <= node.CurrentPoint.X+1; row++ {
		for col := node.CurrentPoint.Y - 1; col <= node.CurrentPoint.Y+1; col++ {
			if !(col == node.CurrentPoint.Y && row == node.CurrentPoint.X) &&
				(row >= 0 && col >= 0) &&
				(len(*graph) > row && len((*graph)[row]) > col) &&
				(*graph)[row][col].NodeType != "o" {
				nodes = append(nodes, (*graph)[row][col])
			}
		}
	}
	return nodes
}

func BuildGraph(matrix [][]string) (*Graph2d, *Node, *Node) {
	var start_node, end_node *Node
	graph := make(Graph2d, len(matrix))
	for i, row := range matrix {
		graph[i] = make([]*Node, len(row))
		for j, el := range row {
			graph[i][j] = NewNode(Point{i, j}, el)
			if el == "s" {
				start_node = graph[i][j]
			} else if el == "e" {
				end_node = graph[i][j]
			}
		}
	}
	return &graph, start_node, end_node
}
