# A star path finder

## Example

```go
func AstarPath() {
	m := NewMap(30, 30)

	// add block grids in middle
	for n := 10; n < 20; n++ {
		m.AddBlock(n, 15)
		m.AddBlock(15, n)
	}

	// add slow grids
	m.AddSlow(1, 1)
	m.AddSlow(2, 2)
	m.AddSlow(2, 3)
	m.AddSlow(3, 3)

	src := m.grids[0][0]
	target := m.grids[29][29]
	pathNode := m.GenPath(src, target)
	if pathNode == nil {
		fmt.Printf("cannot find a path from %v to %v \n", src, target)
		return
	}

	// print map and path
	m.PrintMapWithPath(pathNode)
}
```