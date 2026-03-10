package astar

import (
	"container/heap"
	"fmt"
)

type Map struct {
	grids [][]*Grid
	kind  GridKind
}

func NewQuadMap(width, height int) *Map {
	return newMap(width, height, GridKindQuad)
}

func NewHexMap(width, height int) *Map {
	return newMap(width, height, GridKindHex)
}

func NewGridsMap(xs []int, ys []int, kind GridKind) *Map {
	m := &Map{
		grids: make([][]*Grid, 0),
		kind:  kind,
	}

	newGridFn := NewQuadGrid
	if kind == GridKindHex {
		newGridFn = NewHexGrid
	}

	for idx := range xs {
		x := xs[idx]
		y := ys[idx]

		if m.grids[y] == nil {
			m.grids[y] = make([]*Grid, 0)
		}

		m.grids[y][x] = newGridFn(x, y)
	}

	return m
}

func newMap(width, height int, kind GridKind) *Map {
	if width <= 0 || height <= 0 {
		panic("invalid map width or height")
	}

	m := &Map{
		grids: make([][]*Grid, width),
		kind:  kind,
	}

	for n := 0; n < width; n++ {
		m.grids[n] = make([]*Grid, height)
	}

	newGridFn := NewQuadGrid
	if kind == GridKindHex {
		newGridFn = NewHexGrid
	}

	for y := 0; y < width; y++ {
		for x := 0; x < height; x++ {
			m.grids[y][x] = newGridFn(x, y)
		}
	}

	return m
}

func (m *Map) GridValid(x, y int) bool {
	if y < 0 || y >= len(m.grids) {
		return false
	}
	if m.grids[y] == nil {
		return false
	}
	if x < 0 || x >= len(m.grids[y]) {
		return false
	}

	return m.grids[y][x] != nil
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
	if m.kind == GridKindHex {
		return m.GetHexNeighbors(grid)
	}
	return m.GetQuadNeighbors(grid)
}

func (m *Map) GetQuadNeighbors(grid *Grid) []*Grid {
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

func (m *Map) GetHexNeighbors(grid *Grid) []*Grid {
	hexDirections := [][2]int{
		{1, 0},  // right
		{1, -1}, // top-right
		{0, -1}, // top-left
		{-1, 0}, // left
		{-1, 1}, // bottom-left
		{0, 1},  // bottom-right
	}

	neighbors := make([]*Grid, 0, 6)
	for _, dir := range hexDirections {
		x := grid.X + dir[0]
		y := grid.Y + dir[1]

		if !m.GridValid(x, y) {
			continue
		}

		if m.grids[y][x].Block {
			continue
		}

		neighbors = append(neighbors, m.grids[y][x])
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
