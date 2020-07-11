package main

import "fmt"
func variableZeroValue(){//函数名高亮说明该函数已使用
	var a int
	var s string
	fmt.Println(a,s)
}
func variableInitialValue(){//函数值赋初值
	var a , b int = 3 ,4 //go语言一旦定义就必须要用到，负责就报错
	var s string ="abc"
	fmt.Println(a,b,s)//此处用Printf是需要加上"%d%q",空字符串显示为""
					//用println则不需要，但是空字符串不显示
}
func variableTypeDeduction()  {
	var a,b,c,s  = 3,4,true,"def"
	fmt.Println(a,b,c,s)
}
func variableShorter()  {//省略var，用冒号定义变量,如下，和上面的实现的效果一致
	 a,b,c,s  := 3,4,true,"def"
	fmt.Println(a,b,c,s)
}
func main()  {
	fmt.Println("hello")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}