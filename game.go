package main

import "math/rand"

type point struct{ x, y int }

type game struct {
	width, height int
	snake         []point
	direction     point
	food          point
	score         int
	over, won     bool
	rng           *rand.Rand
}

func newGame(width, height int, rng *rand.Rand) *game {
	g := &game{
		width: width, height: height, rng: rng,
		snake:     []point{{width / 2, height / 2}, {width/2 - 1, height / 2}, {width/2 - 2, height / 2}},
		direction: point{1, 0},
	}
	g.placeFood()
	return g
}

func (g *game) step(input string) {
	if g.over || g.won {
		return
	}
	direction := g.direction
	switch input {
	case "w":
		direction = point{0, -1}
	case "s":
		direction = point{0, 1}
	case "a":
		direction = point{-1, 0}
	case "d":
		direction = point{1, 0}
	case "":
	default:
		return
	}
	if direction.x != -g.direction.x || direction.y != -g.direction.y {
		g.direction = direction
	}
	head := point{g.snake[0].x + g.direction.x, g.snake[0].y + g.direction.y}
	if head.x < 0 || head.x >= g.width || head.y < 0 || head.y >= g.height {
		g.over = true
		return
	}
	eating := head == g.food
	body := g.snake
	if !eating {
		// Хвост освободит клетку на этом же ходу.
		body = body[:len(body)-1]
	}
	for _, cell := range body {
		if head == cell {
			g.over = true
			return
		}
	}
	g.snake = append([]point{head}, g.snake...)
	if eating {
		g.score++
		g.placeFood()
	} else {
		g.snake = g.snake[:len(g.snake)-1]
	}
}

func (g *game) placeFood() {
	occupied := make(map[point]bool, len(g.snake))
	for _, cell := range g.snake {
		occupied[cell] = true
	}
	free := make([]point, 0, g.width*g.height-len(g.snake))
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			cell := point{x, y}
			if !occupied[cell] {
				free = append(free, cell)
			}
		}
	}
	if len(free) == 0 {
		g.won = true
		return
	}
	g.food = free[g.rng.Intn(len(free))]
}
