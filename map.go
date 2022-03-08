package astar

import (
	"container/heap"
	"fmt"
)

type Map struct {
	grids [][]*Grid
}

func NewMap(width, height int) *Map {
	if width <= 0 || height <= 0 {
		panic("invalid map width or height")
	}

	m := &Map{
		grids: make([][]*Grid, width),
	}

	for n := 0; n < width; n++ {
		m.grids[n] = make([]*Grid, height)
	}

	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			m.grids[y][x] = NewGrid(x, y)
		}
	}

	return m
}

func (m *Map) GridValid(x, y int) bool {
	return y >= 0 && y < len(m.grids) && x >= 0 && x < len(m.grids[y])
}

func (m *Map) AddBlock(x, y int) {
	if !m.GridValid(x, y) {
		return
	}

	m.grids[y][x].Block = true
}

func (m *Map) AddSlow(x, y int) {
	if !m.GridValid(x, y) {
		return
	}

	m.grids[y][x].Slow = true
}

func (m *Map) GetNeighbors(grid *Grid) []*Grid {
	neighbors := make([]*Grid, 0, 8)
	for x := grid.X - 1; x <= grid.X+1; x++ {
		for y := grid.Y - 1; y <= grid.Y+1; y++ {
			if x == grid.X && y == grid.Y {
				continue
			}

			if !m.GridValid(x, y) {
				continue
			}

			if m.grids[y][x].Block {
				continue
			}

			neighbors = append(neighbors, m.grids[y][x])
		}
	}

	return neighbors
}

// generate path from src to target
func (m *Map) GenPath(src, target *Grid) *PathNode {
	// init path
	path := NewPath(
		m,
		&PathNode{
			grid: src,
			g:    0, // highest priority
		},
		target,
	)

	// finding a path
	for {
		if path.open.Len() <= 0 {
			break
		}

		curNode := heap.Pop(&path.open).(*PathNode)
		if curNode.grid.Equal(target) {
			return curNode
		}

		curNode.closed = true
		neighbors := path.GetNeighbors(curNode)
		for _, neighborNode := range neighbors {
			if neighborNode.closed {
				continue
			}

			if neighborNode.open {
				newG := curNode.g + LengthValue(curNode.grid, neighborNode.grid)
				if newG < neighborNode.g {
					neighborNode.g = newG
					neighborNode.parent = curNode
					heap.Fix(&path.open, neighborNode.index)
				}
			} else {
				neighborNode.open = true
				neighborNode.parent = curNode
				heap.Push(&path.open, neighborNode)
			}
		}
	}

	return nil
}

func (m *Map) PrintMap() {
	for n := range m.grids {
		fmt.Println()
		for _, grid := range m.grids[n] {
			if grid.Block {
				fmt.Printf(" B ")
			} else if grid.Slow {
				fmt.Printf(" S ")
			} else {
				fmt.Printf(" O ")
			}
		}
		fmt.Println()
	}
}

func (m *Map) PrintMapWithPath(node *PathNode) {
	mapPathNodes := make(map[*Grid]*PathNode)
	for e := node; e != nil; e = e.parent {
		mapPathNodes[e.grid] = e
	}

	for n := range m.grids {
		for _, grid := range m.grids[n] {
			if _, found := mapPathNodes[grid]; found {
				fmt.Printf(" *")
				continue
			}

			if grid.Block {
				fmt.Printf(" B")
			} else if grid.Slow {
				fmt.Printf(" S")
			} else {
				fmt.Printf(" O")
			}
		}
		fmt.Println()
	}
}

func (m *Map) GetPathSections(node *PathNode) []*PathSection {
	sections := make([]*PathSection, 0, 8)

	// length == 0
	if node == nil {
		return sections
	}

	// length == 1
	if node.parent == nil {
		sections = append(sections, &PathSection{
			start: node.grid,
			end:   node.grid,
		})
		return sections
	}

	// length > 2
	prev := node
	curr := node.parent
	prevSlope := GetSlope(prev.grid, curr.grid)
	section := &PathSection{
		start: prev.grid,
	}

	for ; curr != nil; curr = curr.parent {
		slope := GetSlope(prev.grid, curr.grid)
		if slope != prevSlope {
			section.end = prev.grid
			sections = append(sections, section)
			section = &PathSection{
				start: prev.grid,
			}
			prevSlope = slope
		}

		prev = curr
	}

	section.end = prev.grid
	sections = append(sections, section)

	return sections
}
