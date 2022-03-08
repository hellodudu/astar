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
	m.AddSlow(0, 1)
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
	sections := m.GetPathSections(pathNode)
	fmt.Println("get path sections:")
	for _, section := range sections {
		fmt.Printf("%d, %d -> %d, %d\n", section.start.X, section.start.Y, section.end.X, section.end.Y)
	}
}
```

## Result
```
 * * * * * * * O O O O O O O O O O O O O O O O O O O O O O O
 S S O O O O O * O O O O O O O O O O O O O O O O O O O O O O
 O O S O O O O O * O O O O O O O O O O O O O O O O O O O O O
 O O S S O O O O O * O O O O O O O O O O O O O O O O O O O O
 O O O O O O O O O O * O O O O O O O O O O O O O O O O O O O
 O O O O O O O O O O O * O O O O O O O O O O O O O O O O O O
 O O O O O O O O O O O O * O O O O O O O O O O O O O O O O O
 O O O O O O O O O O O O O * O O O O O O O O O O O O O O O O
 O O O O O O O O O O O O O O * O O O O O O O O O O O O O O O
 O O O O O O O O O O O O O O O * O O O O O O O O O O O O O O
 O O O O O O O O O O O O O O O B * O O O O O O O O O O O O O
 O O O O O O O O O O O O O O O B O * O O O O O O O O O O O O
 O O O O O O O O O O O O O O O B O O * O O O O O O O O O O O
 O O O O O O O O O O O O O O O B O O O * O O O O O O O O O O
 O O O O O O O O O O O O O O O B O O O O * O O O O O O O O O
 O O O O O O O O O O B B B B B B B B B B O * O O O O O O O O
 O O O O O O O O O O O O O O O B O O O O O O * O O O O O O O
 O O O O O O O O O O O O O O O B O O O O O O O * O O O O O O
 O O O O O O O O O O O O O O O B O O O O O O O O * O O O O O
 O O O O O O O O O O O O O O O B O O O O O O O O * O O O O O
 O O O O O O O O O O O O O O O O O O O O O O O O O * O O O O
 O O O O O O O O O O O O O O O O O O O O O O O O O * O O O O
 O O O O O O O O O O O O O O O O O O O O O O O O O O * O O O
 O O O O O O O O O O O O O O O O O O O O O O O O O O O * O O
 O O O O O O O O O O O O O O O O O O O O O O O O O O O * O O
 O O O O O O O O O O O O O O O O O O O O O O O O O O O * O O
 O O O O O O O O O O O O O O O O O O O O O O O O O O O O * O
 O O O O O O O O O O O O O O O O O O O O O O O O O O O O * O
 O O O O O O O O O O O O O O O O O O O O O O O O O O O O * O
 O O O O O O O O O O O O O O O O O O O O O O O O O O O O O *
get path sections:
29, 29 -> 28, 28
28, 28 -> 28, 26
28, 26 -> 27, 25
27, 25 -> 27, 23
27, 23 -> 25, 21
25, 21 -> 25, 20
25, 20 -> 24, 19
24, 19 -> 24, 18
24, 18 -> 6, 0
6, 0 -> 0, 0
```