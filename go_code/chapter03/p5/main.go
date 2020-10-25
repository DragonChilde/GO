package main

import "fmt"

func main()  {
	var price float32 = 89.12
	fmt.Println("price=", price)	//price= 89.12

	var num1 float32 = -0.00089
	var num2 float32 = -7809656.09
	fmt.Println("num1=", num1 , "num2=", num2)	//num1= -0.00089 num2= -7.809656e+06

	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3=", num3 , "num4=", num4)	//num3= -123.00009 num4= -123.0000901

	var num5 = 1.1
	fmt.Printf("num5的数据类型是%T", num5)	//num5的数据类型是float64

	num6 := 5.12
	num7 := .123
	fmt.Println("num6=", num6 , "num7=", num7)	//num6= 5.12 num7= 0.123

	num8 := 5.1234e2	//5.1234 * 10的2次方
	num9 := 5.1234E2	//5.1234 * 10的2次方
	num10 := 5.1234E-2	//5.1234 / 10的2次方
	fmt.Println("num8 =", num8 , "num9 =", num9 , "num10 =", num10)	//num8 = 512.34 num9 = 512.34 num10 = 0.051234
}
