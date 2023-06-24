package _case

import "log"

func Fib(n int) int {
	if n <= 2 {
		log.Fatal("请选择大于2的位置")
	}
	t := tool()
	var res int
	for i := 0; i < n-2; i++ {
		res = t()
	}
	return res
}

//斐波那契数列，x1+x2=X2...，求xn的值
func tool() func() int {
	var x0 = 0
	var x1 = 1
	var x2 = 0
	return func() int {
		x2 = x1 + x0
		x0 = x1
		x1 = x2
		return x2
	}
}
