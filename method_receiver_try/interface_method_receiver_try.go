package main

import "fmt"

type TestStruct struct {
	X int
}

// 定义 double 接口
type Doubler interface {
	Double() int
}

// 定义 triple 接口
type Tripler interface {
	Triple() int
}

// 定义在值上的 method
func (ts TestStruct) Double() int {
	r := ts.X * 2
	fmt.Println("ts.Double = ", r)
	return r
}

// 定义在指针上的 method
func (ts *TestStruct) Triple() int {
	r := (*ts).X * 3
	fmt.Println("*ts.Triple = ", r)
	return r
}

// 定义接口
func DoubleFun(d Doubler) {
	d.Double()
}

// 定义接口
func TripleFun(t Tripler) {
	t.Triple()
}

func main() {

	// 接收者是值的方法可以通过指针调用，因为会先根据指针找到值
	t1 := &TestStruct{1}
	DoubleFun(t1)

	// 指针方法可以通过指针调用
	t2 := &TestStruct{2}
	TripleFun(t2)

	// 值方法可以通过值调用
	t3 := TestStruct{3}
	DoubleFun(t3)

	// 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址
	//t4 := TestStruct{4}
	//TripleFun(t4)
	//  报错
	//  cannot use t4 (type TestStruct) as type Tripler in argument to TripleFun:
	//  TestStruct does not implement Tripler (Triple method has pointer receiver)
}