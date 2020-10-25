package main

import (
	_ "fmt"
)

func main()  {

	/*var n1 int32 = 12
	var n2 int64
	var n3 int8*/

	/*n2 = n1 + 20	//int32 ---> int64 错误
	n3 = n1 + 20	//int32 ---> int8 错误*/

/*	n2 = int64(n1) + 20
	n3 = int8(n1) + 20
	fmt.Println("n2 = ", n2 , "n3 = ", n3)	//n2 =  32 n3 =  32
*/

/*	var m1 int32
	var m2 int8
	var m3 int8
	m2 = int8(m1) + 127		//编译通过,但是结果不是127+12,按溢出处理
	m3 = int8(m1) + 128		//编译不通过
	fmt.Println("m2 = ",m2)*/
}
