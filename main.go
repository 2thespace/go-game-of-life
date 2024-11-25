package main

import (
	"fmt"
	"time"
)

type Cell struct {
	isAlived bool
}

type Canvas struct {
	cells          [][]Cell
	size_x, size_y uint16
}

func NewCanvas(x uint16, y uint16) Canvas {
	var tmp Canvas
	tmp.size_x = x
	tmp.size_y = y
	for range x {
		var row []Cell
		for range y {
			var c Cell
			c.isAlived = false
			row = append(row, c)
		}
		tmp.cells = append(tmp.cells, row)
	}
	return tmp
}

func (c Canvas) GetNeibgor(x uint16, y uint16) int {
	if x >= c.size_x || y >= c.size_y {
		return 0
	}
	var neigbor int = 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			pos_x := uint16(int(c.size_x+x)+dx) % c.size_x
			pos_y := uint16(int(c.size_y+y)+dy) % c.size_y
			if c.isAlived(pos_x, pos_y) {
				neigbor += 1
			}
		}
	}
	return neigbor
}

func (c Canvas) isAlived(x uint16, y uint16) bool {
	if x >= c.size_x || y >= c.size_y {
		return false
	}
	return c.cells[x][y].isAlived
}

func printBool(value bool) int {
	if value {
		return 1
	}
	return 0
}

func (c Canvas) printCells() {
	for row := 0; row < int(c.size_x); row++ {
		for column := 0; column < int(c.size_y); column++ {
			fmt.Printf("%d ", printBool(c.cells[row][column].isAlived))
		}
		fmt.Print("\n")
	}
}

func (c Canvas) PutLivedCell(x uint16, y uint16) {
	if x >= c.size_x || y >= c.size_y {
		return
	}
	c.cells[x][y].isAlived = true
}

func (c Canvas) Update() {
	tmp := NewCanvas(c.size_x, c.size_y)
	for x, row := range c.cells {
		for y := range row {
			is_alliving := c.cells[x][y].isAlived
			neigh_count := c.GetNeibgor(uint16(x), uint16(y))
			if neigh_count == 2 && is_alliving || neigh_count == 3 {
				tmp.cells[x][y].isAlived = true
			} else {
				tmp.cells[x][y].isAlived = false
			}

		}
	}
	copy(c.cells, tmp.cells)
	// c.cells = tmp
}

func main() {
	canv := NewCanvas(10, 10)

	canv.PutLivedCell(1, 2)
	canv.PutLivedCell(2, 3)
	canv.PutLivedCell(3, 3)
	canv.PutLivedCell(3, 2)
	canv.PutLivedCell(3, 1)
	for i := 0; i < 70; i++ {
		fmt.Print("\033[0;0H") // move cursor up one line
		canv.printCells()
		canv.Update()
		time.Sleep(500 * time.Millisecond)

	}

}
