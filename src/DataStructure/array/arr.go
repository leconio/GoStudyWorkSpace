package array

import "fmt"

func init() {
	//testInit()
	testCopy()
}

func testCopy() {
	var arr2 [5]string
	arr1 := []string{"red","blue","black","yellow","grey"}
	//arr2 = arr1 不可赋值
	arr2[1] = arr1[1]
	arr2[2] = arr1[2]
	arr2[3] = arr1[3]
	arr2[4] = arr1[4]
	arr2[0] = arr1[0]
	// 以上为值拷贝
	fmt.Println(arr2)
	arr2[0] = "haha"
	fmt.Println(arr1[0])
	printArr(&arr1)
	fmt.Println(&arr1)
}

func printArr(arr *[]string)  {
	//arr[3] = "ddd" FIXME
	fmt.Println(arr)
}

func testInit()  {
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	arr2 := []int{11, 22, 33, 44, 55}
	fmt.Println(arr2)
	//如果使用 ... 替代数组的长度，Go 语言会根据初始化时数组元素的数量来确定该数组的长 度
	arr3 := [...]int{111, 222, 333, 444, 555}
	fmt.Println(arr3)
}