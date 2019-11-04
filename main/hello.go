package main
import (
	"fmt"
	"strconv"
)

//全局变量
var(
	i = 100;
)

func main(){
	//输出
	// helloGo()

	//'\r'使用
	// useSlashR()

	//变量
	// variable()

	//+使用
	// plusUse()

	//string 转为 基本数据类型
	// parseString()

	//指针使用
	aboutPoint()
}

func aboutPoint(){
	var v32 int32 = 123456
	fmt.Printf("v32 = %v \n",v32)
	fmt.Printf("&v32 type %v \n",&v32)
	fmt.Printf("*(&v32) type %v \n",*(&v32))
	var ptr *int32 = &v32
	fmt.Printf("ptr %v \n",ptr)
	fmt.Printf("&ptr %v \n",&ptr)
	fmt.Printf("*ptr %v \n",*ptr)
	*ptr = 1
	fmt.Printf("*ptr %v \n",*ptr)
	fmt.Printf("v32 %v \n",v32)
}



func parseString() {
	v32 := "-354634382"
	var s int64
	s, _ = strconv.ParseInt(v32, 10, 32)
	fmt.Printf("%T, %v\n", s, s)
	fmt.Printf("%T, %v\n", int32(s), int32(s))
	


	v64 := "-3546343826724305832"
	sss,_ := strconv.ParseInt(v64, 10, 64)
	fmt.Printf("%T, %v\n", sss, sss)
	


}
//+使用
func plusUse(){
	fmt.Println("10 + 10.23:",10 + 10.23)
	fmt.Println("'10' + '10.23':","10" + "10.23")
	//fmt.Println("10 + '10.23':",10 + "10.23")
}

//输出
func helloGo(){
	fmt.Println("hello.go","啊手动阀士大夫嘎斯的噶多噶","阿斯顿噶人噶第三个")
}
//'\r'使用
func useSlashR(){
	fmt.Println("hello\rgo 123")
	fmt.Println("hello\rgo s123")
	fmt.Println("hello\rgo f123")
}
//变量
func variable(){
	var val int
	val = 99
	fmt.Println("val=",val)
}