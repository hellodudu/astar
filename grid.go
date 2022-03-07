package astar

type Grid struct {
	X, Y  int  // position
	Block bool // blocked
	Slow  bool // slow down 50% speed
}

func NewGrid(x, y int) *Grid {
	return &Grid{X: x, Y: y}
}

func (g *Grid) Equal(other *Grid) bool {
	return g.X == other.X && g.Y == other.Y
}
