package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	var i int = 1
	fmt.Println("i=", i)

	//int8的范围-128-127
	//其它的int16,int32,int64类推
	var j int8 = 127
	fmt.Println("j=" ,j)


	//uint8的范围0-255,其它的uint16,unit32,nuit64类推
	var k uint8 = 255
	fmt.Println("k=" , k)

	//整型的使用细节
	var n1 = 100	//n1是什么类类型
	//fmt.Printf()可以用于做模式化输出
	fmt.Printf("n1的类型 %T" , n1)


	//如何在程序查看某个变量的占用字节大小和数据类型(使用较多)
	var n2 int64 = 10
	//unsafe.Sizeof(n2) 是unsafe包的一个函数,可以返回n2变量占用的字节数
	fmt.Printf("n2的类型%T n2占用的字节数是 %d", n2 , unsafe.Sizeof(n2))

	//Golang程序中整型变量在使用时,遵守保小不保大的原则
	//即:在保底程序正确运行下,尽量使用占用空间小的数据类型
	var age byte = 90
	fmt.Printf("age的类型%T", age)
}
