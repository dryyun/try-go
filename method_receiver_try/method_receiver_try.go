package main

import "fmt"

type IntReceiver struct {
	X int
}

// 方法接收器声明
func (i IntReceiver) Double() int {
	return i.X * 2
}

// 普通函数声明
func IRDouble(i IntReceiver) int {
	return i.X * 2
}

// 带参数的方法声明
func (i IntReceiver) Add(a int) int {
	return i.X + a
}

// 基于指针的方法声明
func (i *IntReceiver) Triple() (r int) {
	r = (*i).X * 3
	fmt.Println(r)
	return
}

func main() {
	i1 := &IntReceiver{1}
	i1.Triple() // 指针 -》调用 receiver 是指针的 method

	i2 := IntReceiver{2}
	i2.Triple() // 值变量 -》调用 receiver 是指针的 method

	i3 := IntReceiver{3}
	i3.Double() // 值变量 -》调用 receiver 是值的 method

	i4 := &IntReceiver{4}
	i4.Double() // 指针 -》调用 receiver 是值的 method

	//
	IntReceiver{5}.Double() // 正确
	//IntReceiver{6}.Triple() // 错误  cannot take the address of IntReceiver literal
	// 正确做法
	//t := IntReceiver{6}
	//t.Triple()

	(&IntReceiver{7}).Double() // 正确
	(&IntReceiver{8}).Triple() // 正确

}
