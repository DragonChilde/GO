package main

import "fmt"

func main()  {
	//基本数据类型在内存布局
	var i int = 10
	//i的地址是什么,&i
	fmt.Println("i的地址 = " ,&i)		//i的地址 =  0xc00000a0b0

	//1. ptr 是一个指针变量
	//2. ptr的类型 *int
	//3. ptr 本身的值此&i
	var ptr *int = &i
	fmt.Printf("ptr = %v\n" , ptr)	//ptr = 0xc00000a0b0
	fmt.Printf("ptr 的地址 = %v\n" , &ptr)	//ptr 的地址 = 0xc000006030
	fmt.Printf("ptr 指向的值 = %v\n" , *ptr)	//ptr 指向的值 = 10

	var num = 100
	fmt.Printf("num 的地址是 %v\n" , &num)

	var ptr2 *int = &num
	*ptr2 = 200
	fmt.Printf("num 值是 %v\n" , num)


	var a int = 300
	var b int = 400
	var ptr3 *int = &a 	//ok
	*ptr3 = 100 			//a = 100
	ptr3 = &b 			//ok
	*ptr3 = 200			// b = 200
	fmt.Printf("a = %d , b = %d , ptr = %d" , a , b , *ptr3)
}
