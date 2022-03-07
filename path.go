package astar

import (
	"container/heap"
	"math"
)

type PathNode struct {
	grid   *Grid
	parent *PathNode
	g      float64
	h      float64
	index  int
	open   bool
	closed bool
}

type Path struct {
	m      *Map
	root   *PathNode
	nodes  map[*Grid]*PathNode
	open   priorityQueue
	target *Grid
}

func NewPath(m *Map, root *PathNode, target *Grid) *Path {
	path := &Path{
		m:      m,
		root:   root,
		nodes:  make(map[*Grid]*PathNode),
		open:   make([]*PathNode, 0),
		target: target,
	}

	path.nodes[root.grid] = root
	heap.Init(&path.open)

	// add target to open list
	heap.Push(&path.open, path.root)

	return path
}

func (p *Path) GetNeighbors(node *PathNode) []*PathNode {
	neighborNodes := make([]*PathNode, 0)
	neighborGrids := p.m.GetNeighbors(node.grid)
	for _, grid := range neighborGrids {
		neighborNode := p.CheckPathNode(grid)
		if neighborNode.g == 0 {
			gridG := LengthValue(node.grid, neighborNode.grid)
			neighborNode.g = gridG + node.g

			// if grid is slow, double grid's g value
			if neighborNode.grid.Slow {
				neighborNode.g += gridG
			}
		}

		neighborNodes = append(neighborNodes, neighborNode)
	}

	return neighborNodes
}

// if node didn't exist, create one and calculate it's h value
func (p *Path) CheckPathNode(grid *Grid) *PathNode {
	node, found := p.nodes[grid]
	if !found {
		node = &PathNode{
			grid: grid,
			h:    LengthValue(grid, p.target),
		}
		p.nodes[grid] = node
	}

	return node
}

func LengthValue(from, to *Grid) float64 {
	x := from.X - to.X
	y := from.Y - to.Y

	if x == 0 {
		return math.Abs(float64(y))
	}

	if y == 0 {
		return math.Abs(float64(x))
	}

	max := x
	if y > x {
		max = y
	}
	return math.Abs(float64(max)) * math.Sqrt2
	// return math.Sqrt(float64(x*x + y*y))
}
