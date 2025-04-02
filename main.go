package main

import (
	"fmt"
	"log"
	"test/2048/util"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()
	util.ClearScreen()

	fmt.Println("按任意键(Ctrl+C)...")

	g := util.NewGame()
	g.Generate()
	g.PrintGrid()
	vaildKey := true
	moved := false
	for {
		// g.PrintGrid()
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'w':
			moved = g.Move(util.Up)
			vaildKey = true
		case 's':
			moved = g.Move(util.Down)
			vaildKey = true
		case 'a':
			moved = g.Move(util.Left)
			vaildKey = true
		case 'd':
			moved = g.Move(util.Right)
			vaildKey = true
		default:
			vaildKey = false
		}
		if vaildKey && moved {
			if !g.Generate() {
				break
			}
		}
		if g.IsGmaeOver() {
			fmt.Println("Game Over")
		}
		// time.Sleep(1 * time.Second)
		g.PrintGrid()
		fmt.Printf("按键：%c (%d)\n", char, key)
	}
}
