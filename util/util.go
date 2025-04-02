package util

import (
	"fmt"
	"math/rand"
)

const (
	Up = iota
	Down
	Left
	Right
)

const (
	GridSize = 4
)

type Pos struct {
	x int
	y int
}

type Game struct {
	Grid  [GridSize][GridSize]int
	Score int
}

func (g *Game) Move(dir int) bool {
	res := false
	boardSize := len(g.Grid)
	switch dir {
	case Right:
		for i := range boardSize {
			for j := boardSize - 1; j > 0; j-- {
				b := tryMoveNum(&(g.Grid)[i][j-1], &(g.Grid)[i][j])
				if b {
					j = boardSize
					res = true
				}
			}
		}
	case Left:
		for i := range boardSize {
			for j := 0; j < boardSize-1; j++ {
				b := tryMoveNum(&(g.Grid)[i][j+1], &(g.Grid)[i][j])
				if b {
					j = -1
					res = true
				}
			}
		}
	case Down:
		for i := range boardSize {
			for j := boardSize - 1; j > 0; j-- {
				b := tryMoveNum(&(g.Grid)[j-1][i], &(g.Grid)[j][i])
				if b {
					j = boardSize
					res = true
				}
			}
		}
	case Up:
		for i := range boardSize {
			for j := 0; j < boardSize-1; j++ {
				b := tryMoveNum(&(g.Grid)[j+1][i], &(g.Grid)[j][i])
				if b {
					j = -1
					res = true
				}
			}
		}
	}
	return res
}

func (g *Game) findZero() []Pos {
	boardSize := len(g.Grid)
	res := []Pos{}
	for i := range boardSize {
		for j := range boardSize {
			if (g.Grid)[i][j] == 0 {
				res = append(res, Pos{i, j})
			}
		}
	}
	return res
}

func isMoveble(a int, b int) bool {
	if a == 0 && b == 0 {
		return false
	}
	if a == b || b == 0 {
		return true
	}
	return false
}

func Reverse(arr *[]int) {
	for i, j := 0, len(*arr)-1; i < j; i, j = i+1, j-1 {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
}

func tryMoveNum(a *int, b *int) bool {
	if isMoveble(*a, *b) {
		*b = *a + *b
		*a = 0
		return true
	}
	return false
}

func ClearScreen() {
	fmt.Print("\033[2J")
	Repaint()
}

func Repaint() {
	// fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

func (g *Game) PrintGrid() {
	// Repaint()
	res := ""
	for i := range GridSize {
		if i == 0 {
			res += printGridCell(GridSize, "┌", "┬", "┐")
			res += printGridData(GridSize, &(g.Grid)[i])
		} else if i == GridSize-1 {
			res += printGridCell(GridSize, "├", "┼", "┤")
			res += printGridData(GridSize, &(g.Grid)[i])
			res += printGridCell(GridSize, "└", "┴", "┘")
		} else {
			res += printGridCell(GridSize, "├", "┼", "┤")
			res += printGridData(GridSize, &(g.Grid)[i])
		}
	}
	fmt.Print(res)
}

func printGridCell(cellNum int, head string, middle string, end string) string {
	res := head
	cellStr := "────"
	for i := range cellNum {
		if i == cellNum-1 {
			res += cellStr + end
		} else {
			res += cellStr + middle
		}
	}
	return res + "\n"
}

func printGridData(cellNum int, data *[GridSize]int) string {
	res := ""
	for i := range cellNum {
		value := (*data)[i]
		if value == 0 {
			res += "|    "
		} else {
			res += fmt.Sprintf("|%4d", value)
		}
	}
	res += "|"
	return res + "\n"
}

func NewGame() *Game {
	return &Game{
		Grid:  [GridSize][GridSize]int{},
		Score: 0,
	}
}

var seedPool = [2]int{2, 4}

func (g *Game) Generate() bool {
	zeroPos := g.findZero()
	if len(zeroPos) == 0 {
		return false
	}
	index := rand.Intn(len(zeroPos))
	(g.Grid)[zeroPos[index].x][zeroPos[index].y] = seedPool[rand.Intn(len(seedPool))]
	return true
}

func (g *Game) IsGmaeOver() bool {
	zeroPos := g.findZero()
	return len(zeroPos) == 0
}
