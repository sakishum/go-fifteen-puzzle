package main

import (
	"math/rand"
	"time"

	"github.com/rthornton128/goncurses"
)

const (
	tileSize = 3
	rowCount = 4
	colCount = 4
)

func generateNumbers(seed int64) []int {
	numbers := make([]int, rowCount*colCount, rowCount*colCount)

	// Initial state: the last pos is zero
	for i := 1; i <= rowCount*colCount; i++ {
		numbers[i-1] = i % (rowCount * colCount)
	}

	// Randomly swap position
	rand.Seed(seed)
	for i := 0; i < 256; i++ {
		posA := rand.Int31n(rowCount * colCount)
		posB := rand.Int31n(rowCount * colCount)

		numbers[posA], numbers[posB] = numbers[posB], numbers[posA]
	}

	return numbers
}

func findSpacePos(numbers []int) (int, int) {
	for i := 0; i < rowCount; i++ {
		for j := 0; j < colCount; j++ {
			if numbers[i*colCount+j] == 0 {
				return i, j
			}
		}
	}

	panic("error: failed to findSpacePos")
}

func moveDown(numbers []int) {
	row, col := findSpacePos(numbers)

	if row != rowCount-1 {
		numbers[row*colCount+col], numbers[(row+1)*colCount+col] =
			numbers[(row+1)*colCount+col], numbers[row*colCount+col]
	}
}

func moveUp(numbers []int) {
	row, col := findSpacePos(numbers)

	if row != 0 {
		numbers[row*colCount+col], numbers[(row-1)*colCount+col] =
			numbers[(row-1)*colCount+col], numbers[row*colCount+col]
	}
}

func moveLeft(numbers []int) {
	row, col := findSpacePos(numbers)

	if col != 0 {
		numbers[row*colCount+col], numbers[row*colCount+col-1] =
			numbers[row*colCount+col-1], numbers[row*colCount+col]
	}
}

func moveRight(numbers []int) {
	row, col := findSpacePos(numbers)

	if col != colCount-1 {
		numbers[row*colCount+col], numbers[row*colCount+col+1] =
			numbers[row*colCount+col+1], numbers[row*colCount+col]
	}
}

func main() {
	numbers := generateNumbers(time.Now().UnixNano())

	stdscr := initIO()
	defer closeIO()

	drawBackground(stdscr)

	for {
		showStatus(stdscr, solved(numbers))
		drawBoard(stdscr, numbers)
		stdscr.Refresh()

		key := stdscr.GetChar()
		switch key {
		case 'q':
			return
		case 'j', goncurses.KEY_DOWN:
			moveDown(numbers)
		case 'k', goncurses.KEY_UP:
			moveUp(numbers)
		case 'h', goncurses.KEY_LEFT:
			moveLeft(numbers)
		case 'l', goncurses.KEY_RIGHT:
			moveRight(numbers)
		}
	}
}

func solved(numbers []int) bool {
	for i := 1; i <= rowCount*colCount; i++ {
		if numbers[i-1] != i%(rowCount*colCount) {
			return false
		}
	}

	return true
}
