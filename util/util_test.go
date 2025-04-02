package util

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	Reverse(&arr)
	t.Log("Reverse test passed")
	fmt.Println(arr)
}

func TestTryMoveNum(t *testing.T) {
	a := 2
	b := 2
	tryMoveNum(&a, &b)
	fmt.Printf("%d, %d", a, b)
}

func TestMove(t *testing.T) {
	g := &Game{
		Grid:  [GridSize][GridSize]int{{2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 2}, {4, 4, 4, 4}},
		Score: 0,
	}
	// g.Generate()
	g.Move(Down)
	g.PrintGrid()
}

func TestMoveUp(t *testing.T) {
	g2 := &Game{
		Grid:  [GridSize][GridSize]int{{0, 2, 4, 8}, {2, 2, 2, 2}, {2, 2, 2, 2}, {2, 2, 2, 4}},
		Score: 0,
	}
	// g.Generate()
	g2.Move(Up)
	g2.PrintGrid()
}

func TestMoveLeft(t *testing.T) {
	g2 := &Game{
		Grid:  [GridSize][GridSize]int{{0, 2, 4, 8}, {2, 2, 2, 2}, {2, 2, 2, 2}, {4, 2, 2, 2}},
		Score: 0,
	}
	// g.Generate()
	g2.Move(Left)
	g2.PrintGrid()
}

func TestMoveRight(t *testing.T) {
	g2 := &Game{
		Grid:  [GridSize][GridSize]int{{2, 2, 4, 8}, {2, 2, 2, 2}, {2, 2, 2, 2}, {4, 2, 2, 2}},
		Score: 0,
	}
	// g.Generate()
	g2.Move(Right)
	g2.PrintGrid()
}

func TestFindZero(t *testing.T) {
	g := NewGame()
	posArr := g.findZero()
	t.Log(posArr)
}
