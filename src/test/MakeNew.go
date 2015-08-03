package main
import "fmt"

func main() {
	testNew()
	testMake()
}

type Student struct {
	id   int
	name string
}

/**
new用于各种类型的内存分配。内建函数new本质上说跟其它语言中的同名函数功能一样：
new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。
用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。有一点非常重要：Go返回地址
 */
func testNew() {
	//stu1是指针类型，stu2是结构体类型
	stu1 := new(Student)
	stu1.id, stu1.name=1, "yi"
	stu2 := Student{id:2, name:"er"}
	fmt.Println(stu1, stu2)
}
/**
内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型，而不是*T。
本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。
 */
func testMake() {
	slice1 := make([]int, 3)
	fmt.Println(slice1)
	
	map1 := make(map[int]string, 5)
	fmt.Println(map1)
}