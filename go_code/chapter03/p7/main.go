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
