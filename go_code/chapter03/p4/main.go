package main

import "fmt"

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
}
