package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestStep(t *testing.T) {
	tests := []struct {
		name, input string
		snake       []point
		direction   point
		food        point
		wantHead    point
		wantLength  int
		wantScore   int
		wantOver    bool
	}{
		{"forward", "", []point{{3, 3}, {2, 3}, {1, 3}}, point{1, 0}, point{0, 0}, point{4, 3}, 3, 0, false},
		{"turn", "w", []point{{3, 3}, {2, 3}, {1, 3}}, point{1, 0}, point{0, 0}, point{3, 2}, 3, 0, false},
		{"reverse ignored", "a", []point{{3, 3}, {2, 3}, {1, 3}}, point{1, 0}, point{0, 0}, point{4, 3}, 3, 0, false},
		{"eat", "d", []point{{3, 3}, {2, 3}, {1, 3}}, point{1, 0}, point{4, 3}, point{4, 3}, 4, 1, false},
		{"wall", "d", []point{{7, 3}, {6, 3}, {5, 3}}, point{1, 0}, point{0, 0}, point{7, 3}, 3, 0, true},
		{"body", "s", []point{{3, 3}, {2, 3}, {2, 4}, {3, 4}, {4, 4}}, point{1, 0}, point{0, 0}, point{3, 3}, 5, 0, true},
		{"vacated tail", "s", []point{{3, 3}, {2, 3}, {2, 4}, {3, 4}}, point{1, 0}, point{0, 0}, point{3, 4}, 4, 0, false},
		{"unknown input", "hello", []point{{3, 3}, {2, 3}, {1, 3}}, point{1, 0}, point{0, 0}, point{3, 3}, 3, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &game{width: 8, height: 8, snake: tt.snake, direction: tt.direction, food: tt.food, rng: rand.New(rand.NewSource(1))}
			g.step(tt.input)
			if g.snake[0] != tt.wantHead || len(g.snake) != tt.wantLength || g.score != tt.wantScore || g.over != tt.wantOver {
				t.Fatalf("head=%v length=%d score=%d over=%v; want %v, %d, %d, %v", g.snake[0], len(g.snake), g.score, g.over, tt.wantHead, tt.wantLength, tt.wantScore, tt.wantOver)
			}
		})
	}
}

func TestFoodNeverOverlapsSnake(t *testing.T) {
	g := newGame(8, 8, rand.New(rand.NewSource(1)))
	for i := 0; i < 100; i++ {
		g.placeFood()
		if g.food.x < 0 || g.food.x >= g.width || g.food.y < 0 || g.food.y >= g.height {
			t.Fatalf("food outside board: %v", g.food)
		}
		for _, cell := range g.snake {
			if cell == g.food {
				t.Fatalf("food overlaps snake: %v", g.food)
			}
		}
	}
}

func TestEatLastFreeCellWins(t *testing.T) {
	g := &game{
		width: 2, height: 2,
		snake:     []point{{0, 0}, {0, 1}, {1, 1}},
		direction: point{0, -1}, food: point{1, 0},
		rng: rand.New(rand.NewSource(1)),
	}
	g.step("d")
	if !g.won || g.over || g.score != 1 || len(g.snake) != 4 {
		t.Fatalf("expected victory, got %+v", g)
	}
}

func TestFinishedGameDoesNotMove(t *testing.T) {
	for _, won := range []bool{false, true} {
		g := newGame(8, 8, rand.New(rand.NewSource(1)))
		g.won, g.over = won, !won
		before := append([]point(nil), g.snake...)
		g.step("w")
		if !reflect.DeepEqual(before, g.snake) || g.score != 0 {
			t.Fatal("finished game changed")
		}
	}
}
