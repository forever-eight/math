package ds

import "math"

type Point struct {
	X float64
	Y float64
}

type Line struct {
	A float64
	B float64
	C float64
}

func MakeLine(fp Point, sp Point) Line {
	line := Line{}

	line.A = fp.Y - sp.Y           // y1-y2
	line.B = sp.X - fp.X           // x2-x1
	line.C = fp.X*sp.Y - sp.X*fp.Y // x1y2 - x2y1

	return line
}

func (l *Line) MinDistance(p Point) float64 {
	d := math.Abs(l.A*p.X+l.B*p.Y+l.C) / math.Sqrt(l.A*l.A+l.B*l.B)
	return d
}

func (l *Line) Distance(point Point) float64 {
	// y = kx+b
	k := -l.A / l.B
	b := -l.C / l.B
	yLine := k*point.X + b
	d := point.Y - yLine
	return d
}
