package main
import "fmt"

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
	plusUse()
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