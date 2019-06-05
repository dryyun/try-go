package main

import "fmt"

func main() {
	var a1 [3]int                   // 使用 int 零值初始化 a1，a1 就是 [3]int 类型
	fmt.Printf("%T\t%#v\n", a1, a1) // [3]int  [3]int{0, 0, 0}

	a2 := new([3]int)                       // a2 是指针，是 *[3]int类型
	fmt.Printf("%T\t%p\t%#v\n", a2, a2, a2) // *[3]int 0xc0000182c0    &[3]int{0, 0, 0}
	fmt.Println(a2[0], (*a2)[1])            // 0 0 // 都引用了数组的值，其中 a2[0] 是解引用，根据指针找到了变量

	a2[0] = 123
	(*a2)[1] = 234
	fmt.Println(a2, *a2) // &[123 234 0] [123 234 0]

	fmt.Println(len(a2), len(*a2)) // 3 3
}

