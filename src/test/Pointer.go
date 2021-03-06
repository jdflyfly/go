package main

import "fmt"


/**
Go语言中channel，slice，map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。
（注：若函数需改变slice的长度，则仍需要取地址传递指针）
 */
func main(){
	testCallbyValue()
}


func testCallbyValue(){
	x :=3
	add1(x)
	fmt.Println(x)
	add2(&x)
	fmt.Println(x)
}

func add1(x int){
	x=x+1
}

func add2(x *int){
	*x = *x+1
}