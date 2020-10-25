package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var n1 int = 99
	var n2 float64 = 23.456
	var n3 bool = true
	var n4 byte = 'h'
	var str string	//空的str

	//使用第一种方式来转换fmt.Sprintf方法
	str = fmt.Sprintf("%d", n1)
	fmt.Printf("str type %T str=%q\n" , str ,str )

	str = fmt.Sprintf("%f", n2)
	fmt.Printf("str type %T str=%q\n" , str ,str )

	str = fmt.Sprintf("%t", n3)
	fmt.Printf("str type %T str=%q\n" , str ,str )

	str = fmt.Sprintf("%c", n4)
	fmt.Printf("str type %T str=%q\n" , str ,str )


	var n5 int = 99
	var n6 float64 = 23.456
	var n7 bool = true

	str = strconv.FormatInt(int64(n5) , 10)
	fmt.Printf("str type %T str=%q\n" , str ,str )	//str type string str="99"

	//strconv.FormatFloat(n6 , "f" , 10 ,64)
	//'f'格式10 :表示小数位保留10位 64:表示这个小数是float64
	str = strconv.FormatFloat(n6 , 'f',10 ,64)
	fmt.Printf("str type %T str=%q\n" , str ,str )	//str type string str="23.4560000000"

	str = strconv.FormatBool(n7)
	fmt.Printf("str type %T str=%q\n" , str ,str )	//str type string str="true"




	//strconv包中有一个函数Itora
	var num5 int64 = 4567
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T str=%q\n" , str ,str )	//str type string str="4567"


	var str2 string = "true"
	var b bool
	//1. strconv.ParseBool(str) 函数会返回两个值(value bool, err error)
	//2. 只想获取取value bool,不想获取 err 所以使用 _ 忽略
	b , _ = strconv.ParseBool(str2)
	fmt.Printf("b type %T str=%v\n" , b ,b )	//b type bool str=true

	var str3 string = "123456789"
	var m1 int64
	var m2 int
	m1 , _  = strconv.ParseInt(str3 , 10 ,64)
	m2 = int(m1)
	fmt.Printf("m1 type %T str=%v\n" , m1 ,m1 )	//m1 type int64 str=123456789
	fmt.Printf("m2 type %T str=%v\n" , m2 ,m2 )	//m2 type int str=123456789

	var str4 string = "123.456"
	var f2 float64
	f2 ,_  = strconv.ParseFloat(str4 , 64)
	fmt.Printf("f2 type %T str=%v\n" , f2 ,f2 )	//f2 type float64 str=123.456

	var str5 string = "hello"
	var m3 int64 = 11
	var m4 bool = true
	m3 , _ = strconv.ParseInt(str5 , 10 ,64)
	fmt.Printf("m3 type %T str=%v\n" , m3 ,m3 )		//m3 type int64 str=0

	m4 , _  = strconv.ParseBool(str5)
	fmt.Printf("m4 type %T str=%v\n" , m4 ,m4 )		//m4 type bool str=false
}
