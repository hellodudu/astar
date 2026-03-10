package astar

type GridKind int

const (
	GridKindQuad GridKind = iota // 四边形格子
	GridKindHex                  // 六边形格子 (Axial坐标)
)

type Grid struct {
	X, Y  int // position (Axial: X=q, Y=r)
	Kind  GridKind
	Block bool // blocked
	Slow  bool // slow down 50% speed
}

func NewQuadGrid(x, y int) *Grid {
	return &Grid{X: x, Y: y, Kind: GridKindQuad}
}

func NewHexGrid(x, y int) *Grid {
	return &Grid{X: x, Y: y, Kind: GridKindHex}
}

func (g *Grid) Equal(other *Grid) bool {
	return g.X == other.X && g.Y == other.Y
}
