package main

var a int = 10
var b = 10
var c bool

var d, e string = "d", "e"

func main() {
	// 类型相同多个变量, 非全局变量
	var f, g int
	f, g = 10, 100
	println(a, b, c)
	println(d, e)
	println(f, g)
}