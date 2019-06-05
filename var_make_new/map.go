package main

import "fmt"

func main() {
	var m1 map[string]int                   // 初始化为 map 的零值，nil
	fmt.Printf("%T\t%p\t%#v\n", m1, m1, m1) // map[string]int  0x0     map[string]int(nil)
	// 通过 0x0 就可以看出是零值 nil
	fmt.Println(len(m1), m1 == nil) // 0 true

	//m1["a"] = 0 // panic: assignment to entry in nil map
	// 因为 map 是 nil，还没真实分配空间，所以报错
	// 要真正使用，还必须要 m1 = make(map[string]int )
	m1 = make(map[string]int)
	m1["a"] = 0

	fmt.Println("========")

	m2 := map[string]int{}
	fmt.Printf("%T\t%p\t%#v\n", m2, m2, m2) // map[string]int  0xc0000a0000    map[string]int{}
	fmt.Println(len(m2), m2 == nil)         // 0 false

	m2["a"] = 0 // 这步没有报错，是因为已经分配空间了

	fmt.Println("========")

	m3 := make(map[string]int)

	fmt.Printf("%T\t%p\t%#v\n", m3, m3, m3)
	fmt.Println(len(m3), m3 == nil) // 0 false

	// m2，m3 声明其实是一样的，行为一样
	// 包括 m4 的声明都差不都
	//m4 := map[string]int{
	//	"a": 1,
	//}

	fmt.Println("========")

	m5 := new(map[string]int)
	fmt.Printf("%T\t%p\t%#v\n", m5, m5, m5)    // *map[string]int 0xc00000e010    &map[string]int(nil)
	fmt.Printf("%T\t%p\t%#v\n", *m5, *m5, *m5) // map[string]int  0x0     map[string]int(nil)

	// m5 是指向 map 的指针
	// *m5 是 nil 的 map

	*m5 = make(map[string]int)
	(*m5)["a"] = 1
	fmt.Printf("%T\t%p\t%#v\n", m5, m5, m5)    // *map[string]int 0xc00000e010    &map[string]int{"a":1}
	fmt.Printf("%T\t%p\t%#v\n", *m5, *m5, *m5) // map[string]int  0xc000068270    map[string]int{"a":1}

	// 要使用 *m5，必须先 make


}
