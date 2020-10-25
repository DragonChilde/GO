package main

import "fmt"

func main()  {
	var address string = "广州天河"
	fmt.Println("address is ", address)

	//字符串一旦赋值了,字符串就不能修改了,在GO中字符串是不可变的
	//var str = "hello"
	//str[0] = "a"		//这里就不能去修改str的内容,即go中的字符串是不可变的

	str2 := "abc\nabc"
	fmt.Println("str2 is " , str2)

	//使用的反引号``
	str3 := `package main

		import "fmt"
		
		func main()  {
			var address string = "广州天河"
			fmt.Println("address is ", address)
		}`
	fmt.Println("str3 is ", str3)

	//字符串拼接方式
	str := "hello" + "world"
	fmt.Println("str is ", str)

	//当一个拼接的操作很长时,可以分行写,但要注意,需要将+保留在上一行(go默认每行会自动补全分号;)
	str4 := "hello" + "world" + "hello" + "world" + "hello" + "world" + "hello" +
		"hello" + "world" + "hello" + "world" + "hello" + "world" +
		"hello" + "world" + "hello" + "world"
	fmt.Println("str3 is ", str4)


	var a int	//0
	var b float32	//0
	var c float64	//0
	var d bool		//
	var e string	// ""
	fmt.Printf("a = %d , b = %v , c =%v ,d = %v , e = %v" ,a , b ,c , d, e)



}
