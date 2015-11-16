package main

import "fmt"

func main() {
	//testStruct()
	testStructEqual()
}

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human  // 匿名字段，struct
	Skills // 匿名字段，自定义的类型string slice
	int    // 内置类型作为匿名字段
	speciality string
}

func testStruct() {
	// 初始化学生Jane
	jane := Student{Human:Human{"Jane", 35, 100}, speciality:"Biology"}
	//其他初始化工作
	//human := Human{name:"Jane", age:35, weight:100}
	//jane := Student{Human:human, speciality:"java"}

	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her speciality is ", jane.speciality)
	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int)
}

/**
Struct values are comparable if all their fields are comparable. Two
    struct values are equal if their corresponding non-blank
    <http://golang.org/ref/spec#Blank_identifier> fields are equal.
 */

type Obj struct {
	i int
	s string
	_ int //blank identifier
}


func testStructEqual() {
	obj1 := Obj{i:1, s:"s"}
	obj2 := Obj{i:1, s:"s"}
	p1 := &obj1
	p2 := &obj2

	fmt.Printf("obj1:%v,obj2:%v,obj1==obj2?%v\n", obj1, obj2, obj1==obj2)
	fmt.Printf("p1:%v,p2:%v,p1==p2?%v\n", p1, p2, p1==p2)


}