# linux下安装GO

1. 解压文件到/opt/下

   ```bash
   tar -zxvf go1.15.3.linux-amd64.tar.gz -C /opt/
   ```

2. linux下配置Golang环境变量

   ```bash
   vim /etc/profilevim /etc/profile
   ```

3. 参数如下

   ```bash
   export GOROOT=/opt/go
   export PATH=$PATH:$GOROOT/bin
   export GOPATH=$HOME/goproject
   ```

4. 刷新环境变量

   ```bash
   source /etc/profile
   ```


# Go语言开发入门

## 开发步骤

1. 开发程序项目时,go的目录结构处理

   ​	![](http://120.77.237.175:9080/photos/go/01.jpg)

   1. `go`文件的后缀是.go
   2. `package main` 表示该 hello.go 文件所在的包是 main, 在 go 中，每个文件都必须归属于一个包。
   3. `import "fmt"` 表示：引入一个包，包名 fmt, 引入该包后，就可以使用 fmt 包的函数，比如：fmt.Println
   4. `fun main(){}` `func` 是一个关键字，表示一个函数。`main` 是函数名，是一个主函数，即我们程序的入口。
   
2. 通过 ·go build· 命令对该 go 文件进行编译，生成 .exe 文件.       运行生成的main.exe 文件即可

   > 注意：通过 go run 命令可以直接运行 hello.go 程序 [类似执行一个脚本文件的形式]

   

## Golang 执行流程分析

  如果是对源码编译后，再执行，Go 的执行流程如下图

![](http://120.77.237.175:9080/photos/go/02.jpg)

  如果我们是对源码直接 执行 go run 源码，Go 的执行流程如下图

![](http://120.77.237.175:9080/photos/go/03.jpg)

>   两种执行流程的方式区别
>
> ​    如果我们先编译生成了可执行文件，那么我们可以将该可执行文件拷贝到没有 go 开发环境的机 器上，仍然可以运行
>
> ​	如果我们是直接 go run go 源代码，那么如果要在另外一个机器上这么运行，也需要 go 开发 环境，否则无法执行。
>
> ​    在编译时，编译器会将程序运行依赖的库文件包含在可执行文件中，所以，可执行文件变大了很多。

## 编译和运行说明

 有了 go 源文件，通过编译器将其编译成机器可以识别的二进制码文件。在该源文件目录下，通过 go build 对 hello.go 文件进行编译。可以指定生成的可执行文件名，在 windows 下 必须是 .exe 后缀。

```go
go build -o 指定生成的文件名 要编译的文件名.go
```

 如果程序没有错误，没有任何提示，会在当前目录下会出现一个可执行文件(**windows 下是.exe Linux 下是一个可执行文件**)，该文件是二进制码文件，也是可以执行的程序。

如果程序有错误，编译时，会在错误的那行报错。有助于程序员调试错误.

## Go 程序开发的注意事项

1. Go 源文件以 "go" 为扩展名。
2. Go 应用程序的执行入口是 main()函数。 这个是和其它编程语言（比如 java/c）
3. Go 语言严格区分大小写。
4. Go 方法由一条条语句构成，每个语句后不需要分号(Go 语言会在每行后自动加分号)，这也体现出 Golang 的简洁性。
5. Go 编译器是一行行进行编译的，因此我们一行就写一条语句，不能把多条语句写在同一个，否则报错
6. Go 语言定义的变量或者 **import** 的包如果没有使用到，代码不能编译通过。
7. 大括号都是成对出现的，缺一不可。

## Go 语言的转义字符(escape char)

说明:常用的转义字符有如下:

1. \t : 表示一个制表符，通常使用它可以排版。

2. \n  ：换行符

3.  \\  ：一个\

4. \"  ：一个"

5. 5) \r    ：一个回车   fmt.Println("天龙八部雪山飞狐\r 张飞"); 

   ![](http://120.77.237.175:9080/photos/go/04.jpg)


使用一句输出语句,达到以下效果

![](http://120.77.237.175:9080/photos/go/05.jpg)

##       Golang 开发常见问题和解决方法

![](http://120.77.237.175:9080/photos/go/06.jpg)

##       注释(comment)

### 在 Golang 中注释有两种形式

```go
	//行注释
	/*块注释*/
```

##       规范的代码风格

1. 正确的注释和注释风格： Go 官方推荐使用行注释来注释整个方法和语句。

2. 正确的缩进和空白

   1. 使用一次 **tab** 操作，实现缩进,默认整体向右边移动，时候用 **shift+tab** 整体向左移

   2.  或者使用 gofmt **来进行格式化**

      ```go
      //该指令可以将格式后的内容重写入到文件.当重写打开该文件时,就会看到新的格式化后的文件
      gofmt -w xxx.go
      ```

3. 运算符两边习惯性各加一个空格。比如：2 + 4 * 5。

   ```go
   var num = 2 + 4 * 5
   ```

4. Go 语言的代码风格.

   ```go
   package main
   import "fmt"
   //注意:{}前第一个{括号不能首开一行,必须跟在方法后(GO语法规定)
   func main()  {
   	fmt.Println("Hello world!")
   }
   ```

5.  一行最长不超过 80 个字符，超过的请使用换行展示，尽量保持格式优雅

# Golang变量

##       变量的介绍

### 变量的概念

变量相当于内存中一个数据存储空间的表示，你可以把变量看做是一个房间的门牌号，通过门牌号我们可以找到房间，同样的道理，通过变量名可以访问到变量(值)。

### 变量的使用步骤

      1. 声明变量(也叫:定义变量)
         2. 非变量赋值
         3. 使用变量

```go
package main
import "fmt"
func main()  {
	//定义变量 /声明变量
	var i int
	//给i赋值
	i = 10
    //使用变量
	fmt.Println(i)
}
```

##           变量使用注意事项

1. 变量表示内存中的一个存储区域

2. 该区域有自己的名称（变量名）和类型（数据类型）

3. Golang 变量使用的三种方式

   1. 第一种：指定变量类型，声明后若不赋值，使用默认值

      ```go
      	/*
      	golang的变量使用方式一
      	第一种:指定变量类型,声明后若不赋值,使用默认值
      	int的默认值是0
      	 */
      	var i int
      	fmt.Println("i =" ,i)
      ```

      

   2. 根据值自行判定变量类型(类型推导)

      ```go
      	/*
      	第二种:根据值自行判定变量类型(类型推导)
      	 */
      	var num = 1.1
      	fmt.Println("num =", num)
      ```

      

   3. 省略 var, 注意 :=左侧的变量不应该是已经声明过的，否则会导致编译错误

      ```go
      /*
      第三种:省略var,注意:=左侧的变量不应该是已经声明的,否则会导致编译错误
      下面的方式等价 var name string name = "tom"
       */
      name := "tom"
      fmt.Println("name=", name)
      ```

   4. 多变量声明

      在编程中，有时我们需要一次性声明多个变量，Golang 也提供这样的语法

      ```go
      package main
      import "fmt"
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
      }
      ```
      
      如何一次性声明多个全局变量【在 go 中函数外部定义变量就是全局变量】：

      ```go
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
      	fmt.Println("m1=", m1 , "name3=", name3 ,"m2=", m2)
      	fmt.Println("m3=", m3 , "name4=", name4 ,"m4=", m4)
      }
      ```
   
   
      5. 该区域的数据值可以在同一类型范围内不断变化(重点)
   

​    ![](http://120.77.237.175:9080/photos/go/07.jpg)

6. 变量在同一个作用域(在一个函数或者在代码块)内不能重名
   
   ​    ![](http://120.77.237.175:9080/photos/go/08.jpg)

>    变量=变量名+值+数据类型，这一点请大家注意，变量的三要素 Golang 的变量如果没有赋初值，编译器会使用默认值, 比如 int 默认值 0  string 默认值为空串，小数默认为 0

##       变量的声明，初始化和赋值

![](http://120.77.237.175:9080/photos/go/09.jpg)

##       程序中 +号的使用

1. 当左右两边都是数值型时，则做加法运算
2.  当左右两边都是字符串，则做字符串拼接

## 数据类型的基本介绍

![](http://120.77.237.175:9080/photos/go/10.jpg)

##       整数类型

![](http://120.77.237.175:9080/photos/go/11.jpg)

![](http://120.77.237.175:9080/photos/go/12.jpg)

```go
var i int = 1
fmt.Println("i=", i)

//int8的范围-128-127
//其它的int16,int32,int64类推
var j int8 = 127
fmt.Println("j=" ,j)

//uint8的范围0-255,其它的uint16,unit32,nuit64类推
var k uint8 = 255
fmt.Println("k=" , k)
```

![](http://120.77.237.175:9080/photos/go/13.jpg)

### 整型的使用细节

1. Golang 各整数类型分：有符号和无符号，int uint 的大小和系统有关。

2. Golang 的整型默认声明为 int 型

   ```go
   //整型的使用细节
   var n1 = 100	//n1是什么类类型
   //fmt.Printf()可以用于做模式化输出
   fmt.Printf("n1的类型 %T" , n1)	//n1的类型 int
   ```

3. 如何在程序查看某个变量的字节大小和数据类型 （使用较多）

   ```go
   //如何在程序查看某个变量的占用字节大小和数据类型(使用较多)
   var n2 int64 = 10
   //unsafe.Sizeof(n2) 是unsafe包的一个函数,可以返回n2变量占用的字节数
   fmt.Printf("n2的类型%T n2占用的字节数是 %d", n2 , unsafe.Sizeof(n2))
   ```

4. `Golang` 程序中整型变量在使用时，遵守保小不保大的原则，即：在保证程序正确运行下，尽量 使用占用空间小的数据类型。【如：年龄】

   ```go
   //Golang程序中整型变量在使用时,遵守保小不保大的原则
   //即:在保底程序正确运行下,尽量使用占用空间小的数据类型
   var age byte = 90
   ```

   > bit: 计算机中的最小存储单位。byte:计算机中基本存储单元。[二进制再详细说] 1byte = 8 bit

##       小数类型/浮点型

```go
var price float32 = 89.12
fmt.Println("price=", price)	//price= 89.12
```

### 小数类型分类

![](http://120.77.237.175:9080/photos/go/14.jpg)

1. 关于浮点数在机器中存放形式的简单说明,浮点数=符号位+指数位+尾数位 说明：浮点数都是有符号的.

   ```go 
   var num1 float32 = -0.00089
   var num2 float32 = -7809656.09
   fmt.Println("num1=", num1 , "num2=", num2)	//num1= -0.00089 num2= -7.809656e+06
   ```

2. 尾数部分可能丢失，造成精度损失

   ``` go
   var num3 float32 = -123.0000901
   var num4 float64 = -123.0000901
   fmt.Println("num3=", num3 , "num4=", num4)	//num3= -123.00009 num4= -123.0000901
   ```

   > 说明：float64 的精度比 float32 的要准确.如果我们要保存一个精度高的数，则应该选用 float64

3. 浮点型的存储分为三部分：符号位+指数位+尾数位 在存储过程中，精度会有丢失

### 浮点型使用细节

1. `Golan` 浮点类型有固定的范围和字段长度，不受具体 OS(操作系统)的影响。

2. `Golang` 的浮点型默认声明为 float64 类型。

   ```go
   var num5 = 1.1
   fmt.Printf("num5的数据类型是%T", num5)	//num5的数据类型是float64
   ```

3. 浮点型常量有两种表示形式

   1. 十进制数形式：如：5.12和.512    (必须有小数点）

      ```go
      num6 := 5.12
      num7 := .123
      fmt.Println("num6=", num6 , "num7=", num7)	//num6= 5.12 num7= 0.123
      ```

   2. 科学计数法形式:如：5.1234e2   = 5.12 * 10 的 2 次方     5.12E-2   = 5.12/10 的 2 次方

      ```go
      num8 := 5.1234e2	//5.1234 * 10的2次方
      num9 := 5.1234E2	//5.1234 * 10的2次方
      num10 := 5.1234E-2	//5.1234 / 10的2次方
      fmt.Println("num8=", num8 , "num9=", num9 , "num10=", num10)	
      //num8 = 512.34 num9 = 512.34 num10 = 0.051234
      ```

      > 通常情况下，应该使用 float64 ，因为它比 float32 更精确。[开发中，推荐使用 **float64**]

##       字符类型

Golang 中没有专门的字符类型，如果要存储单个字符(字母)，一般使用 **byte** 来保存。字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。也就是说对于传统的字符串是由字符组成的，而 **Go** 的字符串不同，它是由字节组成的。

```go
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
}
```

      1. 如果保存的字符在 `ASCII` 表的,比如`[0-1, a-z,A-Z..]`直接可以保存到 `byte`
         2. 如果保存的字符对应码值大于 255,这时我们可以考虑使用 `int` 类型保存
         3. 如果需要安装字符的方式输出，这时我们需要格式化输出，即 `fmt.Printf(“%c”, c1)`..

###  字符类型使用细节

1. 字符常量是用单引号('')括起来的单个字符。例如：`var c1 byte = 'a' var c2 int = '中' var c3 byte = '9'`

2. Go 中允许使用转义字符 '\’来将其后的字符转变为特殊字符型常量。例如：`var c3 char = ‘\n’ '\n'`表示换行符

3. Go 语 言 的 字 符 使 用 UTF-8 编 码 ， 如 果 想 查 询 字 符 对 应 的 utf8 码 值 http://www.mytju.com/classcode/tools/encode_utf8.asp **英文字母-1 个字节  汉字-3 个字节**

4. 在 Go 中，字符的本质是一个整数，直接输出时，是该字符对应的 UTF-8 编码的码值。

5. 可以直接给某个变量赋一个数字，然后按格式化输出时%c，会输出该数字对应的 unicode 字符

   ```go
   //可以直接给某个变量赋一个数字,然后按格式化输出时%c,会输出该数字对应的unicode字符
   var c4 int = 22269
   fmt.Printf("c4 = %c" , c4)	//c4 = 国
   ```

6. 字符类型是可以进行运算的，相当于一个整数，因为它都对应有 Unicode 码.

   ```go
   //字符类型是可以执行运算的,相当于一个整数,输出时是按照码值运行
   var n1 = 10 + 'a'
   fmt.Println("n1 = ", n1)		//n1 =  107
   ```

### 字符类型本质探讨

1. 字符型 存储到 计算机中，需要将字符对应的码值（整数）找出来 
   - 存储：字符--->对应码值---->二进制-->存储 
   - 读取：二进制----> 码值 ----> 字符 --> 读取
2. 字符和码值的对应关系是通过字符编码表决定的(是规定好)
3. Go 语言的编码都统一成了 `utf-8`。非常的方便，很统一，再也没有编码乱码的困扰了

##   布尔类型

1. 布尔类型也叫 bool 类型，bool 类型数据只允许取值 true 和 false
2. bool 类型占 1 个字节。
3. bool 类型适于逻辑运算，一般用于程序流程控制
   - if 条件控制语句
   - for 循环控制语句

```go
package main
import (
	"fmt"
	"unsafe"
)
func main()  {
	var b =false
	fmt.Println("b = ", b)
	//注意
	//1.bool类型占用存储空间是1个字节
	fmt.Println("b的占用空间是", unsafe.Sizeof(b))
	//2. bool类型只能取true或者false
}
```

## String 类型

字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由**单个字节**连接起来的。Go语言的字符串的字节使用 **UTF-8 编码**标识 Unicode 文本

```go
var address string = "广州天河"
fmt.Println("address is ", address)
```

###         String 使用注意事项和细节

1. Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本，这样 Golang 统一使用 UTF-8 编码,中文 乱码问题不会再困扰程序员。

2. 字符串一旦赋值了，字符串就不能修改了：在 Go 中字符串是不可变的。

   ```go
   //字符串一旦赋值了,字符串就不能修改了,在GO中字符串是不可变的
   //var str = "hello"
   //str[0] = "a"		//这里就不能去修改str的内容,即go中的字符串是不可变的
   ```

3. 字符串的两种表示形式

   1. 双引号, 会识别转义字符

   ```go
   str2 := "abc\nabc"
   fmt.Println("str2 is " , str2)
   ```

   2. 反引号，以字符串的原生形式输出，包括换行和特殊字符，可以实现防止攻击、输出源代码等效果 

   ```go
   //使用的反引号``
   str3 := `package main
   
   import "fmt"
   
   func main()  {
   var address string = "广州天河"
   fmt.Println("address is ", address)
   }`
   fmt.Println("str3 is ", str3)
   ```

   3. 字符串拼接方式

      ```go
      //字符串拼接方式
      str := "hello" + "world"
      fmt.Println("str is ", str)
      ```

   4. 当一行字符串太长时，需要使用到多行字符串，可以如下处理

      ```go
      //当一个拼接的操作很长时,可以分行写,但要注意,需要将+保留在上一行	(go默认每行会自动补全分号;)
      str4 := "hello" + "world" + "hello" + "world" + "hello" + "world" + "hello" +
      		"hello" + "world" + "hello" + "world" + "hello" + "world" +
      		"hello" + "world" + "hello" + "world"
      	fmt.Println("str3 is ", str4)
      ```

   

##       基本数据类型的默认值

在 go 中，数据类型都有一个默认值，当程序员没有赋值时，就会保留默认值，在 go 中，默认值又叫零值。

![](http://120.77.237.175:9080/photos/go/15.jpg)

```go
var a int	//0
var b float32	//0
var c float64	//0
var d bool		//
var e string	// ""
fmt.Printf("a = %d , b = %v , c =%v ,d = %v , e = %v" ,a , b ,c , d, e)	//a = 0 , b = 0 , c =0 ,d = false , e = 
```

##       基本数据类型的相互转换

Golang 和 java / c 不同，Go 在不同类型的变量之间赋值时需要显式转换。也就是说 Golang 中数据类型不能自动转换。

### **基本语法**

表达式 T(v) 将值 v 转换为类型 T

- **T**: 就是数据类型，比如 int32，int64，float32 等等
- **v**: 就是需要转换的变量

```go
var i int32 = 100
//将i => float
var n1 float32 = float32(i)
var n2 int8 = int8(i)
var n3 int64 = int64(i)		//低精度=>高精度
fmt.Println("i = %v" , i , "n1 = ", n1 , "n2 = ", n2, "n3 = ", n3)	//i = %v 100 n1 =  100 n2 =  100 n3 =  100
```

### 基本数据类型相互转换的注意事项

   1. Go 中，数据类型的转换可以是从 **表示范围小-->表示范围大，也可以 范围大--->范围小**

   2. 被转换的是变量存储的数据(即值)，**变量本身的数据类型并没有变化**！

      ```go
      //被转换的是变量存储的数据(即会上),变量本身的数据类型并没有变化
      fmt.Printf("i type is  %T", i)		//i type is  int32
      ```

   3. 在转换中，比如将 int64      转成 int8 【-128---127】 ，编译时不会报错，只是转换的结果是按溢出处理，和我们希望的结果不一样。 因此在转换时，需要考虑范围

      ```go
      //在转换中,比如将int64转换成 int8 [-128 --- 127] ,编译时不会报错
      //只是转换的结果是按溢出处理,和所期待的结果不一样
      var num1 int64 = 99999
      var num2 int8 = int8(num1)
      fmt.Println("num2 = ", num2)	//num2 =  -97
      ```

      **测试**

      ```go
      var n1 int32 = 12
      var n2 int64
      var n3 int8
      
      /*n2 = n1 + 20	//int32 ---> int64 错误
      	n3 = n1 + 20	//int32 ---> int8 错误*/
      
      n2 = int64(n1) + 20
      n3 = int8(n1) + 20
      fmt.Println("n2 = ", n2 , "n3 = ", n3)		//n2 =  32 n3 =  32
      ```

      ```go
      var m1 int32
      var m2 int8
      var m3 int8
      m2 = int8(m1) + 127		//编译通过,但是结果不是127+12,按溢出处理
      m3 = int8(m1) + 128		//编译不通过
      ```

      > 注意:在日常开发中,如果`import`包在没有使用的情况下又想删除但又没用到想保留的话可以使用如下格式

```go
import (
	_ "fmt"
)
```

###  基本数据类型和 string 的转换

在程序开发中，经常将基本数据类型转成 string,或者将 string 转成基本数据类型。

1. fmt.Sprintf("%参数", 表达式) 【个人习惯这个，灵活】 

   ```go
   func Sprintf(format string, a ...interface{}) string
   //Sprintf根据format参数生成格式化的字符串并返回该字符串。
   ```

   参数需要和表达式的数据类型相匹配`fmt.Sprintf()..` 会返回转换后的字符串

   ```go
   var n1 int = 99
   var n2 float64 = 23.456
   var n3 bool = true
   var n4 byte = 'h'
   var str string	//空的str
   
   //使用第一种方式来转换fmt.Sprintf方法
   str = fmt.Sprintf("%d", n1)
   fmt.Printf("str type %T str=%q\n" , str ,str )	//str type string str="99"
   
   str = fmt.Sprintf("%f", n2)
   fmt.Printf("str type %T str=%q\n" , str ,str )	//str type string str="23.456000"
   
   str = fmt.Sprintf("%t", n3)
   fmt.Printf("str type %T str=%q\n" , str ,str )	//str type string str="true"
   
   str = fmt.Sprintf("%c", n4)
   fmt.Printf("str type %T str=%q\n" , str ,str )	//str type string str="h"
   ```

2. 使用 strconv 包的函数

   ```go
   func ParseFloat(s string, bitSize int) (f float64, err error)
   func FormatBool(b bool) string
   func FormatInt(i int64, base int) string
   func FormatUint(i uint64, base int) string
   func Itoa(i int) string
   ```
```
   
   ```go
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
```

###   String 类型转基本数据类型

```GO
func ParseBool(str string) (value bool, err error)
func ParseInt(s string, base int, bitSize int) (i int64, err error)
func ParseUint(s string, base int, bitSize int) (n uint64, err error)
func ParseFloat(s string, bitSize int) (f float64, err error)
```

```go
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
```

> 注意: 因为返回的是`int64`或者`float64`,如希望得到`int32,float32.....`等其它的类型,可进行如下处理

```go
var num5 int32
num5 = int32(num)
```

### String 转基本数据类型的注意事项

在将 String 类型转成 基本数据类型时，要确保 **String** 类型能够转成有效的数据，比如 我们可以把 "123" , 转成一个整数，但是不能把 "hello" 转成一个整数，如果这样做，Golang 直接将其转成 0 ，其它类型也是一样的道理. float => 0 bool => false

```go
var str5 string = "hello"
var m3 int64 = 11
m3 , _ = strconv.ParseInt(str5 , 10 ,64)
fmt.Printf("m3 type %T str=%v\n" , m3 ,m3 )		//m3 type int64 str=0(没有转成功,默认值为0)

var m4 bool = true
m4 , _  = strconv.ParseBool(str5)
fmt.Printf("m4 type %T str=%v\n" , m4 ,m4 )		//m4 type bool str=false(没有转成功,默认值为false)
```

## 指针

1. 基本数据类型，变量存的就是值，也叫值类型

2. 获取变量的地址，用&，比如： `var num int`, 获取 `num` 的地址：`&num `

   分析一下基本数据类型在内存的布局.

 	![](http://120.77.237.175:9080/photos/go/16.jpg)

3. 指针类型，指针变量存的是一个地址，这个地址指向的空间存的才是值 比如：`var ptr *int = &num`

    举例说明：指针在内存的布局.

   ![](http://120.77.237.175:9080/photos/go/17.jpg)

4. 获取指针类型所指向的值，使用：`*`，比如：`var ptr *int`, 使用`*ptr` 获取 `ptr` 指向的值

   ```go
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
   }
   ```

5. 案例说明 

![](http://120.77.237.175:9080/photos/go/18.jpg)

**案例**

1. 获取一个 int 变量 num 的地址，并显示到终端

2.  将 num 的地址赋给指针 ptr2 , 并通过 ptr2 去修改 num 的值.

   ```go
   var num = 100
   fmt.Printf("num 的地址是 %v\n" , &num)	//num 的地址是 0xc0000a2090
   
   var ptr2 *int = &num
   *ptr2 = 200		//修改时,会到num的值修改
   fmt.Printf("num 值是 %v\n" , num)	//num 值是 200
   ```


练习

```go
func main(){
	var a int = 300
	var ptr *int = a	//错误(a是整型,ptr是指针,无法将整型赋值到指针)
}

func main(){
    var a int = 300
    var ptr *float32 = &a	//错误(a是int,ptr是float32类型的指针,类型不匹配)
}

func main(){
    var a int = 300
    var b int = 400
    var ptr *int = &a 	//ok
    *ptr = 100 			//a = 100
    ptr = &b 			//ok
    *ptr = 200			// b = 200
    fmt.Ptrinf('a = %d , b = %d , ptr = &d' , a , b , *ptr)		//a = 100 , b = 200 , ptr = 200
    
}
```

### 指针的使用细节

1. 值类型，都有对应的指针类型， 形式为 ***数据类型**，比如 int 的对应的指针就是 *int, float32 对应的指针类型就是 *float32, 依次类推。
2. 值类型包括：基本数据类型 **int** 系列**, float** 系列**, bool, string** 、**数组**和**结构体 struct**

## 值类型和引用类型

### 值类型和引用类型的说明

1. 值类型：基本数据类型 int 系列, float 系列, bool, string 、数组和结构体 struct

2. 引用类型：指针、slice 切片、map、管道 chan、interface 等都是引用类型

###  值类型和引用类型的使用特点

1. 值类型：变量直接存储值，内存通常在栈中分配 

![](http://120.77.237.175:9080/photos/go/19.jpg)

2. 引用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据(值)，内存通常在堆 上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，由 GC 来回收

   ![](http://120.77.237.175:9080/photos/go/20.jpg)

3.  内存的栈区和堆区示意图

   ![](http://120.77.237.175:9080/photos/go/21.jpg)

## 标识符的命名规范

### 标识符概念

​      `Golang` 对各种变量、方法、函数等命名时使用的字符序列称为标识符,凡是自己可以起名字的地方都叫标识符

### 标识符的命名规则

1. 由 26 个英文字母大小写，0-9 ，_ 组成

2. 数字不可以开头。

   ```go
   var num int 	//ok    
   var 3num int 	//error
   ```

3. `Golang`中严格区分大小

   ```go
   var num int = 10
   var Num int = 20
   fmt.Printf("num is %d Num is %d" , num , Num)	//num is 10 Num is 20 在 golang 中，num 和 Num  是两个不同的变量
   ```

4. 标识符不能包含空格

   ```go
   var ab c int = 30	//标只符不能包含空格
   ```

5. 下划线"_"本身在 `Go` 中是一个特殊的标识符，称为空标识符。可以代表任何其它的标识符，但 是它对应的值会被忽略(比如：忽略某个返回值)。所以仅能被作为占位符使用，不能作为标识符使用

   ```go
   var _ int = 40	//error		_是空标识符,用于占用
   ```

6. 不能以系统保留关键字作为标识符（一共有 25 个），比如 break，if 等等..

   ```go
   var hello string		//ok
   var hello12 string		//ok
   var 1hello string		//error	不能以数字开头
   var h-b	string			//error 不能使用-
   var x h string			//error 不能使用空格
   var h_4 string 			//ok
   var _ab string			//ok
   var int int				//ok 强烈建议不要这样使用
   var float32 float32		//ok 强烈建议不要这样使用
   var _ string			//error
   var Abc string			//ok
   ```

### 标识符命名注意事项

1. 包名：保持 package 的名字和目录保持一致，尽量采取有意义的包名，简短，有意义，不要和标准库不要冲突 `fmt`
   ![](http://120.77.237.175:9080/photos/go/24.jpg)
2. 变量名、函数名、常量名：采用驼峰法

   ```GO
   var stuName string = "tom"	//形式:xxxYyyyyZzzz
   var goodPrice float32 = 123.4
   ```

3. 如果变量名、函数名、常量名首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用 ( 注：可以简单的理解成，**首字母大写是公开的，首字母小写是私有的**) ,在 `golang` 没有public , private 等关键字
  ![](http://120.77.237.175:9080/photos/go/22.jpg)

  ![](http://120.77.237.175:9080/photos/go/23.jpg)

## 系统保留关键字

![](http://120.77.237.175:9080/photos/go/25.jpg)

## 系统的预定义标识符

![](http://120.77.237.175:9080/photos/go/26.jpg)