package slice

import "fmt"

func init() {
	fmt.Println("Slice======================================")
	var slice1 = make([]string, 3)
	slice0 := []string{}
	fmt.Println(slice0)
	fmt.Println(len(slice0))
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	// 其长度为 3 个元素，容量为 5 个元素
	var slice2 = make([]string, 3, 5)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	slice2[0] = "lecon0"
	slice2[1] = "lecon1"
	slice2[2] = "lecon2"
	//slice2[3] = "lecon3"   panic: runtime error: index out of ran
	slice2 = append(slice2, "lecon3")
	slice2 = append(slice2, "lecon4")
	slice2 = append(slice2, "lecon5")
	fmt.Println(slice2)
	testSlice1()
	testSlice2()
	testSlice3()
	testSlice4()
}

func testSlice4() {
	slice1 := []int{10, 20, 30, 40, 50}
	// 传递的是指针和容量
	change(slice1)
	fmt.Println(slice1)
}

func change(slice []int) {
	slice[3] = 80
}

func testSlice3() {
	slice1 := []int{10, 20, 30, 40, 50}
	for index, value := range slice1 {
		fmt.Println(index,value)
	}
}

func testSlice1() {
	slice1 := []int{10, 20, 30, 40, 50}
	slice2 := slice1[1:3]
	fmt.Println(slice2) // [20 30]
	slice2[0] = 60
	fmt.Println(slice1) //[10 60 30 40 50]
}

func testSlice2() {
	slice1 := []int{10, 20, 30, 40, 50}
	slice2 := slice1[1:3]
	// 注意：通过Append 覆盖了原数组的元素
	slice2 = append(slice2, 80)
	fmt.Println(slice2) //[20 30 80]
	fmt.Println(slice1) //[10 20 30 80 50]
}
