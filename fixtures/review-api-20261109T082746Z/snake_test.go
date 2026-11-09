package snake

import "testing"

func TestRight(t *testing.T) {
    got := Move(Point{}, "right")
    if got.X != 1 { t.Fatalf("unexpected point: %+v", got) }
}
