package _case

import "fmt"

func SimpleCase() {
	var a, b = 3, 4
	var c, d float64 = 5, 6
	fmt.Println(sumInt(a, b))
	fmt.Println(sumFloat(c, d))
	fmt.Println(sum(a, b))
	fmt.Println(sum(c, d))
}

func sumInt(a, b int) int {
	return a + b
}

func sumFloat(a, b float64) float64 {
	return a + b
}

func sum[T int | float64](a, b T) T {
	return a + b
}
