package main

import (
	"fmt"

	"exer1/polygons"
)

func PrintPerimeter(t polygons.Triangle) {
	fmt.Print("Периметр равен = ")
	fmt.Println(t.GetPerimeter())
}

func PrintSquare(t polygons.Triangle) {
	fmt.Print("Площадь равна = ")
	fmt.Println(t.GetSquare())
}

func inputTriangle() polygons.Triangle {
	var a, b, c int
	fmt.Println("Введите длину одной стороны")
	fmt.Scan(&a)
	fmt.Println("Введите длину второй стороны")
	fmt.Scan(&b)
	fmt.Println("Введите длину третьей стороны")
	fmt.Scan(&c)
	return polygons.MakeTriangle(float32(a), float32(b), float32(c))
}

func chekRectangular(t polygons.Triangle) {
	if t.IsRectangular() {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Не прямоугольный")
	}

}
func main() {
	t := inputTriangle()
	PrintPerimeter(t)
	PrintSquare(t)
	chekRectangular(t)
}
