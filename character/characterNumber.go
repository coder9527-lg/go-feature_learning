package character

import "fmt"

type simple struct {
	value int
}

func TestCharacterNumber() {
	a := simple{
		value: 10,
	}
	//通用占位符
	fmt.Printf("默认格式的值：%v \n", a)
	fmt.Printf("包含字段名的值：%+v \n", a)
	fmt.Printf("go语法表示的值：%#v \n", a)
	fmt.Printf("go语法表示的类型：%T \n", a)
	fmt.Printf("输出字面上的百分号：%%10 \n")
}
