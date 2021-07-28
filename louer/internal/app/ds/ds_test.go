package ds

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLine(t *testing.T) {
	// arrange
	A := Point{
		X: 0,
		Y: 0,
	}

	B := Point{
		X: 1,
		Y: 1,
	}

	// act
	l := MakeLine(A, B)

	// assert
	assert.Equal(t, float64(-1), l.A)
	assert.Equal(t, float64(1), l.B)
	assert.Equal(t, float64(0), l.C)
}

// проверяет разницу оординат между точкой под прямой
func TestDiff_Below(t *testing.T) {
	// arrange
	//	l: проходит через (0, 0); (1, 1)
	A := Point{
		X: 0,
		Y: 0,
	}
	B := Point{
		X: 1,
		Y: 1,
	}
	l := MakeLine(A, B)
	//	точка замера (2, 0)
	C := Point{
		X: 2,
		Y: 0,
	}
	// act
	d := l.Distance(C)
	// assert
	assert.Equal(t, float64(-2), d)
	//	разница оординат между точкой и прямой: -2, это значит что точка под прямой
}

// проверяет разницу оординат между точкой на прямой
func TestDiff_Zero(t *testing.T) {
	// arrange
	//	l: проходит через (0, 0); (1, 1)
	A := Point{
		X: 0,
		Y: 0,
	}
	B := Point{
		X: 1,
		Y: 1,
	}
	l := MakeLine(A, B)
	fmt.Println(l.A, l.B, l.C)
	//	точка замера (3, 3)
	C := Point{
		X: 6,
		Y: 6,
	}
	// act
	d := l.Distance(C)
	// assert
	assert.Equal(t, float64(0), d)
	//	разница оординат между точкой и прямой: 0, это значит что точка лежит на прямой
}

// проверяет разницу оординат между точкой под прямой
func TestDiff_Above(t *testing.T) {
	// arrange
	//	l: проходит через (0, 0); (1, 1)
	A := Point{
		X: 0,
		Y: 0,
	}
	B := Point{
		X: 1,
		Y: 1,
	}
	l := MakeLine(A, B)
	fmt.Println(l.A, l.B, l.C)
	//	точка замера (4, 6)
	C := Point{
		X: 4,
		Y: 6,
	}
	// act
	d := l.Distance(C)
	// assert
	assert.Equal(t, float64(2), d)

	//	разница оординат между точкой и прямой: 2, это значит что точка над прямой
}
