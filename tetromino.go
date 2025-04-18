package main

import "fmt"

type Point struct {
	x, y int
}

type Tetromino struct {
	blocks []Point    
	label  byte
}

func newTetromino(pattern []string, label byte) (*Tetromino, error) {
	

	blocks := make([]Point, 0, 4)
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if pattern[x][y] == '#' {
				blocks = append(blocks, Point{x, y})
			}
		}
	}

	if len(blocks) != 4 {
		return nil, fmt.Errorf("invalid number of blocks")
	}

	if !isValidTetromino(blocks) {
		return nil, fmt.Errorf("invalid tetromino shape")
	}

	// Normalize position to top-left corner
	minX, minY := blocks[0].x, blocks[0].y
	for _, p := range blocks {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}

	}
	

	for i := range blocks {
		blocks[i].x -= minX
		blocks[i].y -= minY

	}
	

	return &Tetromino{blocks: blocks, label: label}, nil
}

func isValidTetromino(blocks []Point) bool {
	

	// Check connectivity using flood fill
	connected := make(map[Point]bool)
	connected[blocks[0]] = true
	change:=true


	for change{
		change=false
		for _, block := range blocks {
			if !connected[block] {
			  continue
				
			}

			// Check adjacent blocks
			adjacent := []Point{
				{block.x + 1, block.y},
				{block.x - 1, block.y},
				{block.x, block.y + 1},
				{block.x, block.y - 1},
			}

			for _, adj := range adjacent {
				if !connected[adj] {
					for _, b := range blocks {
						if b.x == adj.x && b.y == adj.y {
							connected[adj] = true
							change=true
						
						}
					}
				}
			}
		}
	}
		
	

	return len(connected) == 4
}
