package equations

import (
	"fmt"
	"math/cmplx"
)

func SquareEquation() {
	var inputed float32
	fmt.Println("Введите число")
	fmt.Scan(&inputed)
	a := complex(inputed, 0)
	fmt.Println("Введите число")
	fmt.Scan(&inputed)
	b := complex(inputed, 0)
	fmt.Println("Введите число")
	fmt.Scan(&inputed)
	c := complex(inputed, 0)
	a128 := complex128(a)
	b128 := complex128(b)
	var x complex128
	if a == 0 {
		if b == 0 {
			fmt.Println("Некорректный ввод")
		} else {
			fmt.Println("X равен ")
			x = complex128(-c / b)
		}
	} else {
		if b == 0 {
			fmt.Println("X равен  ")
			x = complex128(cmplx.Sqrt(complex128(-c / a)))
		} else {
			if c == 0 {
				fmt.Println("X равен ")
				x = complex128(-b / a)
			} else {
				d := complex128(b*b - 4*a*c)
				d = cmplx.Sqrt(d)

				if d == 0 {
					fmt.Println("X равен ")
					x = ((-d) / 2) / a128
				} else {
					fmt.Println("X1 равен ")
					x = ((0 - b128 - d) / 2) / a128
					fmt.Println(x)
					fmt.Println("X2 равен ")
					x = ((0 - b128 + d) / 2) / a128
				}

			}
		}
	}
	fmt.Println(x)

}
