package main
import "fmt"
func main(){
	helloGo()
	// useSlashR()
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
