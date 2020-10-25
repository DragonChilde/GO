package main

import "fmt"

func main()  {
	var i int32 = 100
	//将i => float
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Println("i = %v" , i , "n1 = ", n1 , "n2 = ", n2, "n3 = ", n3)	//i = %v 100 n1 =  100 n2 =  100 n3 =  100

	//被转换的是变量存储的数据(即会上),变量本身的数据类型并没有变化
	fmt.Printf("i type is  %T", i)		//i type is  int32

	//在转换中,比如将int64转换成 int8 [-128 --- 127] ,编译时不会报错
	//只是转换的结果是按溢出处理,和所期待的结果不一样
	var num1 int64 = 99999
	var num2 int8 = int8(num1)
	fmt.Println("num2 = ", num2)	//num2 =  -97
}