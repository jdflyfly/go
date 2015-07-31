package main

import "fmt"

func main() {
	TestArray()

}

func TestArray() {
	var a = []int{1, 2, 3}
	fmt.Println(a)
	fmt.Println(a==nil)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	s :="中国人"
	b :=[]rune(s)
	b[0]='b'
	s2 :=string(b)

	fmt.Println(b)
	fmt.Println(s)
	fmt.Println(s2)
}
