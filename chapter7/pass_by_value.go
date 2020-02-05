package main

import (
	"fmt"
	"unsafe"
)

//64位操作系统
type ZeroStructTest struct {
}

//32 byte len
type ContentStructTest struct {
	n int //8 byte
	s string //16 byte
	p *int //8 byte
}

//
// string type:
// type stringStruct struct {
//  	str unsafe.Pointer //8 byte
//      len int            //8 byte
// }
//

//go 语言是值传递

func stringFunc(s string) {
	fmt.Printf("%p\n", &s)
	fmt.Println("string: ", unsafe.Sizeof(s))
}

func arrayFunc(arr [4]int) {
	fmt.Printf("%p\n", &arr)
	fmt.Println("[4]int: ", unsafe.Sizeof(arr))
}


//
// slice structure:
// type slice struct {
//		array unsafe.Pointer // 8 bytes
//		len   int            // 8 bytes
//		cap   int            // 8 bytes
//	}
//

func sliceFunc(sl []int) {
	fmt.Printf("%p\n", &sl)
	fmt.Println("[]int: ", unsafe.Sizeof(sl))
}

// 指定传地址
func slicePointerFunc(sl *[]int) {
	fmt.Printf("%p\n", sl)
	fmt.Println("*[]int: ", unsafe.Sizeof(sl))
}

func pointerFunc(i *int) {
	fmt.Printf("%p\n", i)
	fmt.Println("*int: ", unsafe.Sizeof(i))
}

// zero struct
func zeroStructFunc(zs ZeroStructTest) {
	fmt.Printf("%p\n", &zs)
	fmt.Println("zeroStructTest: ", unsafe.Sizeof(zs))
}

// content struct
func contentStructFunc(cs ContentStructTest) {
	fmt.Printf("%p\n", &cs)
	fmt.Println("contentStructTest: ", unsafe.Sizeof(cs))
}

func main() {
	fmt.Println("=====String=====")
	s := "str str\\0 str"
	fmt.Println(s)
	fmt.Printf("%p\n", &s)
	stringFunc(s)

	fmt.Println("=====Array=====")
	arr := [4]int{1,2,3,4}
	fmt.Printf("%p\n", &arr)
	arrayFunc(arr)

	sl := make([]int, 2)
	fmt.Println("=====Slice=====")
	fmt.Printf("%p\n", &sl)
	sliceFunc(sl)
	fmt.Printf("%p\n", &sl)

	i := 100
	fmt.Println("=====Pointer=====")
	fmt.Printf("%p\n", &i)
	pointerFunc(&i)
	slicePointerFunc(&sl)

	zs := ZeroStructTest{}
	fmt.Println("=====ZeroStruct=====")
	fmt.Printf("%p\n", &zs)
	zeroStructFunc(zs)

	ecs := ContentStructTest{}
	fmt.Println("=====Empty ContentStruct=====")
	fmt.Printf("%p\n", &ecs)
	contentStructFunc(ecs)

	cs := ContentStructTest{1, "changwang", &i}
	fmt.Println("=====Full ContentStruct=====")
	fmt.Printf("%p\n", &cs)
	contentStructFunc(cs)
}
