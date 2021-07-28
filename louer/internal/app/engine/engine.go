package engine

import (
	"fmt"
	"time"

	"github.com/forever-eight/trade-math/internal/app/ds"
)

func Process() {
	A := ds.Point{
		X: float64(time.Now().AddDate(0, 0, -1).Unix()),
		Y: 0,
	}
	B := ds.Point{
		X: float64(time.Now().Unix()),
		Y: 5.7,
	}
	C := ds.Point{
		X: 1,
		Y: 5,
	}
	line := ds.MakeLine(A, B)
	fmt.Println(line.Distance(C))
}
