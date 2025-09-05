package polygons

import (
	"fmt"
	"math"
)

type Triangle struct {
	LengthA float32
	LengthB float32
	LengthC float32
}

func try_catch() {

	if r := recover(); r != nil {
		fmt.Println("Ошибка:", r)
	}
}

func MakeTriangle(a float32, b float32, c float32) Triangle {
	t := Triangle{
		LengthA: 0,
		LengthB: 0,
		LengthC: 0,
	}
	defer try_catch()
	if (a > 0) && (b > 0) && (c > 0) {
		if ((a + b) > c) && ((a + c) > b) && ((c + b) > a) {
			t = Triangle{
				LengthA: a,
				LengthB: b,
				LengthC: c,
			}
		} else {
			panic("Одна из сторон равно или больше двух других")
		}
	} else {
		panic("Сторона или стороны длиной равно или меньше 0")
	}
	return t
}

func (t Triangle) GetPerimeter() float32 {
	return t.LengthA + t.LengthB + t.LengthC
}

func (t Triangle) GetSquare() float64 {
	p := t.GetPerimeter() / 2

	return (math.Sqrt(float64(((p - t.LengthA) * (p - t.LengthB) * (p - t.LengthC) * p))))

}

func (t Triangle) IsRectangular() bool {
	a2 := t.LengthA * t.LengthA
	b2 := t.LengthB * t.LengthB
	c2 := t.LengthC * t.LengthC

	return (a2+b2 == c2) || (c2+b2 == a2) || (a2+c2 == b2)
}
