package main

import "fmt"

func main() {
	var s1 []int                            // 使用 slice 零值就是 nil 初始化 s1
	fmt.Printf("%T\t%p\t%#v\n", s1, s1, s1) // []int   0x0     []int(nil)
	// s1 的地址是 0x0，就是说还没有分配地址
	fmt.Println(len(s1), s1 == nil) // 0 true，长度是 0，s1 = nil

	fmt.Println("========")

	s2 := []int{}   // s2 是指针，可以获取地址
	fmt.Printf("%T\t%p\t%#v\n", s2, s2, s2) // []int   0x118fff0       []int{}
	fmt.Println(len(s2), s2 == nil)         // 0 false

	//s1[0] = 1 // panic: runtime error: index out of range
	//s2[0] = 2 // panic: runtime error: index out of range
	// 因为 s1,s2 len 、cap 都是 0 ，所以赋值会报错


	// 对比 s1，s2 可以发现，s1 长度是 0 ，本身等于 nil，s2 长度也是 0 ，不等于 nil，
	// s1,s2 不同的声明方式，有差异

	fmt.Println("========")

	s3 := make([]int, 2)   // 或者  = make([]int,2,4) 加上第三个 cap 参数都可以
	fmt.Printf("%T\t%p\t%#v\n", s3, s3, s3)  // []int   0xc0000ae000    []int{0, 0}
	fmt.Println(len(s3), cap(s3), s3 == nil) // 2  2  false

	// s3 是指针，len = cap = 2 , 底层是数组

	fmt.Println("========")

	s4 := new([]int)   // s4 是指向开辟的匿名 slice 的指针，而匿名的 slice 使用零值也就是 nil 初始化
	fmt.Printf("%T\t%p\t%T\t%p\n", s4, s4, *s4, *s4) // *[]int  0xc0000a2060    []int   0x0
	// 看结果可以，s4 本身地址有值，*s4 就是切片了，0x0 意味着就是 nil

	fmt.Println(len(*s4), *s4 == nil) // 0 true

	// 怎么使用了，其实 *s4 就是切片了，当做一般变量使用就行了
	*s4 = append(*s4, 1, 2, 3)
	fmt.Printf("%T\t%p\t%T\t%p\n", s4, s4, *s4, *s4) // *[]int  0xc0000a2060    []int   0xc0000a0020
	fmt.Println(len(*s4), cap(*s4), *s4 == nil)      // 3 4 false

	*s4 = append(*s4, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("%T\t%p\t%T\t%p\n", s4, s4, *s4, *s4) // *[]int  0xc0000a2060    []int   0xc0000aa000
	fmt.Println(len(*s4), cap(*s4), *s4 == nil)      // 14 14 false

	// 通过 append 操作发现
	// *s4 的指针地址会扩容，会改变，len 、cap 都会变
	// s4 的地址没有改变

}
