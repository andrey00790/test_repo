package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	g := newGame(20, 12, rand.New(rand.NewSource(time.Now().UnixNano())))
	input := bufio.NewScanner(os.Stdin)
	for {
		render(g)
		if g.over {
			fmt.Println("Game over!")
			return
		}
		if g.won {
			fmt.Println("You win!")
			return
		}
		fmt.Print("Move [W/A/S/D + Enter], Enter = forward, Q = quit: ")
		if !input.Scan() {
			if err := input.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "read input:", err)
				os.Exit(1)
			}
			return
		}
		command := strings.ToLower(strings.TrimSpace(input.Text()))
		if command == "q" {
			return
		}
		g.step(command)
	}
}

func render(g *game) {
	board := make([][]byte, g.height)
	for y := range board {
		board[y] = []byte(strings.Repeat(" ", g.width))
	}
	if !g.won {
		board[g.food.y][g.food.x] = '*'
	}
	for _, cell := range g.snake {
		board[cell.y][cell.x] = 'o'
	}
	board[g.snake[0].y][g.snake[0].x] = '@'
	fmt.Printf("\nSnake | Score: %d\n", g.score)
	fmt.Println("+" + strings.Repeat("-", g.width) + "+")
	for _, row := range board {
		fmt.Printf("|%s|\n", row)
	}
	fmt.Println("+" + strings.Repeat("-", g.width) + "+")
}
