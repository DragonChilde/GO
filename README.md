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