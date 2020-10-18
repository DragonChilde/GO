package main

import "fmt"
//定义全局变量
var m1 = 100
var m2 = 200
var name3 = "kary"
//上面的声明方式,也可以改成一次性声明
var (
	m3 = 300
	m4 = 500
	name4 = "lucy"
)

func main()  {
	//如何一次声明多个变量
	var n1 , n2 , n3 int
	fmt.Println("n1=", n1, "n2=", n2 ,"n3=", n3)

	//一次性声明多个变量的方式2
	var j1 , name , j3 = 100 , "jack" , 999
	fmt.Println("j1=", j1 , "name=", name ,"j3=", j3)

	//一次性声明多个变量的方式3,同样可以使用类型推导
	k1, name2 , k2 := 200 , "tom" , 888
	fmt.Println("k1=", k1 , "name2=", name2 ,"k2=", k2)

	fmt.Println("m1=", m1 , "name3=", name3 ,"m2=", m2)

	fmt.Println("m3=", m3 , "name4=", name4 ,"m4=", m4)
}
