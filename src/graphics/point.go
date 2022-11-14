package graphics

type Point struct {
	X int
	Y int
}

func MarshalPoint(x int, y int) Point {
	return Point{X: x, Y: y}
}
