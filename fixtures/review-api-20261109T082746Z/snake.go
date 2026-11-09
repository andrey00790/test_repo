package snake

type Point struct { X, Y int }

func Move(head Point, direction string) Point {
    next := head
    switch direction {
    case "up": next.Y++
    case "down": next.Y--
    case "left": next.X--
    case "right": next.X++
    }
    return next
}
