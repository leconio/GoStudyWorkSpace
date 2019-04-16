package _struct

import "fmt"

/*
内置类型是由语言提供的一组类型。我们已经见过这些类型，分别是数值类型、字符串类型 和布尔类型。这些类型本质上是原始的类型。因此，当对这些值进
行增加或者删除的时候，会创 建一个新值。

Go 语言里的引用类型有如下几个：切片、映射、通道、接口和函数类型。当声明上述类型 的变量时，创建的变量被称作标头（header）值。从技术细节上说，
字符串也是一种引用类型。 每个引用类型创建的标头值是包含一个指向底层数据结构的指针。每个引用类型还包含一组独特 的字段，用于管理底层数据结构。
因为标头值是为复制而设计的，所以永远不需要共享一个引用 类型的值。标头值里包含一个指针，因此通过复制来传递一个引用类型的值的副本，本质上就是
在共享底层数据结构。

结构类型可以用来描述一组数据值， 这组值的本质即可以是原始的，也可以是非原始的。如 果决定在某些东西需要删除或者添加某个结构类型的值时该结构
类型的值不应该被更改，那么需 要遵守之前提到的内置类型和引用类型的规范。
当需要修改值本身时，在程序中其他地方，需要使用 指针来共享这个值。
 */

type user struct {
	name  string
	age   int
	email string
}

type admin struct {
	person user
	level  string
}
type Duration int64


func init() {
	user1 := user{"lecon", 154, "lecon@lecon.io"}
	fmt.Println(user1)
	user2 := user{
		name:  "spawn",
		age:   19,
		email: "leconio@outlook.com",
	}
	fmt.Println(user2)

	admin1 := admin{
		user{
			"liucl", 26, "edqwd",
		},
		"top",
	}
	fmt.Println(admin1)

	var time1 Duration
	//time1 = int64(1000)   cannot use int64(1000) (type int64) as type Duration in assignment
	fmt.Println(time1)

}
