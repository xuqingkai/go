package main

import "strconv"

func main() {
	info := myInfo()
	name, age := myInfo1()
	info = "My name is " + name + ",I'm " + strconv.Itoa(age) + " years old"
	print(info)
}
func myInfo() string {
	var name string = "zhangsan"
	var age int = 22
	info := "My name is " + name + ",I'm " + strconv.Itoa(age) + " years old"
	return info
}
func myInfo1() (string, int) {
	var name string = "zhangsan"
	var age int = 22
	name = name + "feng"
	age = 22 + 10
	return name, age
}
func getMyInfo(name string, age int) int {
	println("My name is ", name)
	println("I'm ", age)
	info := 100
	return info
}
