// 程序所属包 package只能出现在文件开头 
package main

// 导包
import "fmt"

// 常量定义 string可省略
const NAME string = "imooc"

// 变量定义 string可省略
var mainName string = "main name"

// 一般类型声明
type imoocInt int

// 结构体声明
type Learn struct {

}

// 声明接口
type Ilearn interface {

}

// 函数定义
func learnImooc() {
	fmt.Print("learnImooc func")
}

// main函数
func main()  {
	fmt.Println("learn1")
	fmt.Print(mainName)
	fmt.Println(NAME)
	learnImooc()
}
