//用于演示golang中实现自定义的引用类型
package main

import "fmt"

type Data struct {
	name string
}

type MyRef *Data

func MakeMyRef() MyRef {
	var r MyRef = new(Data)
	return r
}

func change(r MyRef) {
	r.name = "world"
}

func main() {
	var r1 = MakeMyRef()
	r1.name = "hello"
	change(r1)
	fmt.Println(r1.name)
}
