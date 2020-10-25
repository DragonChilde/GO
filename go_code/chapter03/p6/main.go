package main
import "fmt"
func main()  {
	var c1 byte = 'a'
	var c2 byte = '0'	//字符的0

	//当直接输出byte值,就是输出了对应的字符的码值
	fmt.Println("c1 = ", c1)	//c1 =  97
	fmt.Println("c2 = ", c2)	//c2 =  48

	//输出对应的字符,需要使用格式化输出
	fmt.Printf("c1 = %c c2 = %c" , c1 , c2)	//c1 = a c2 = 0

	//var c3 byte = '北'	//constant 21271 overflows byte 溢出
	var c3 int = '北'
	fmt.Printf("c3 = %c c3对应的码值=%d" , c3 , c3)	//c3 = 北 c3对应的码值=21271

	//可以直接给某个变量赋一个数字,然后按格式化输出时%c,会输出该数字对应的unicode字符
	var c4 int = 22269
	fmt.Printf("c4 = %c" , c4)	//c4 = 国

	//字符类型是可以执行运算的,相当于一个整数,输出时是按照码值运行
	var n1 = 10 + 'a'
	fmt.Println("n1 = ", n1)		//n1 =  107
}
