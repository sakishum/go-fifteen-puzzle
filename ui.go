package main

import (
	"log"
	"strconv"

	"github.com/rthornton128/goncurses"
)

func initIO() *goncurses.Window {
	stdscr, err := goncurses.Init()
	if err != nil {
		log.Fatal("init", err)
	}

	goncurses.StartColor()
	goncurses.Cursor(0)
	goncurses.Echo(false)
	stdscr.Keypad(true)
	return stdscr
}

func closeIO() {
	goncurses.End()
}

func showStatus(window *goncurses.Window, solved bool) {
	window.Standend()
	rows, cols := window.MaxYX()
	width := tileSize*2*4 + 3*2
	height := tileSize*4 + 3
	topRow := (rows - height) / 2
	topCol := (cols - width) / 2

	status := ""
	if solved {
		status = "YOU WIN"
	} else {
		status = "playing"
	}

	window.MovePrintf(topRow-2, topCol+width/2-3, status)
}

func drawBackground(window *goncurses.Window) {
	window.Standend()
	rows, cols := window.MaxYX()
	width := tileSize*2*4 + 3*2
	height := tileSize*4 + 3
	topRow := (rows - height) / 2
	topCol := (cols - width) / 2

	window.MovePrintf(topRow+height+1, topCol-2,
		"HELP:")
	window.MovePrintf(topRow+height+2, topCol-2,
		"    h: Left j: Down k: Up l: Right")
	window.MovePrintf(topRow+height+3, topCol-2,
		"    q: Quit")
}

func drawTile(window *goncurses.Window, startRow, startCol, number int) {
	goncurses.InitPair(2, goncurses.C_BLACK, goncurses.C_CYAN)
	window.AttrSet(goncurses.ColorPair(2))
	for i := startRow; i < startRow+tileSize; i++ {
		for j := startCol; j < startCol+tileSize*2; j++ {
			window.Move(i, j)
			window.AddChar(' ')
		}
	}
	window.MovePrintf(startRow+tileSize/2, startCol+tileSize-1, strconv.Itoa(number))
}

func hideTile(window *goncurses.Window, startRow, startCol, number int) {
	goncurses.InitPair(3, goncurses.C_BLACK, goncurses.C_BLACK)
	window.AttrSet(goncurses.ColorPair(3))
	for i := startRow; i < startRow+tileSize; i++ {
		for j := startCol; j < startCol+tileSize*2; j++ {
			window.Move(i, j)
			window.AddChar(' ')
		}
	}
	window.MovePrintf(startRow+tileSize/2, startCol+tileSize-1, strconv.Itoa(number))
}

func drawBoard(window *goncurses.Window, numbers []int) {
	window.Standend()
	rows, cols := window.MaxYX()
	width := tileSize*2*colCount + 3*2
	height := tileSize*rowCount + 3
	topRow := (rows - height) / 2
	topCol := (cols - width) / 2

	// top border
	for i := 0; i <= width+2; i += 2 {
		window.MovePrintf(topRow-1, topCol+i-2, "<>")
	}

	// bottom border
	for i := 0; i <= width; i += 2 {
		window.MovePrintf(topRow+height, topCol+i-2, "<>")
	}

	// left border
	for i := 0; i <= height; i++ {
		window.MovePrintf(topRow+i, topCol-2, "<>")
	}

	// right border
	for i := 0; i <= height; i++ {
		window.MovePrintf(topRow+i, topCol+width, "<>")
	}

	// draw tile
	for x := 0; x < rowCount; x++ {
		for y := 0; y < colCount; y++ {
			number := numbers[x*colCount+y]
			if number != 0 {
				drawTile(window, topRow+tileSize*x+x, topCol+tileSize*y*2+y*2, number)
			} else {
				hideTile(window, topRow+tileSize*x+x, topCol+tileSize*y*2+y*2, number)
			}
		}
	}
}
