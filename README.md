#0 安装
可以访问<html src="https://golang.org">https://golang.org</html>下载安装
#1 环境变量配置
windows下需要配置GOROOT(go语言安装目录),GOPATH(go程序目录)并在path中添加%GOROOT%/bin
暂时还不知道为什么添加GOPATH?
#2 go执行
```
go build xx.go //编译go源代码 => 得到xx.exe文件
xx.exe //执行

go run xx.go //直接执行
```
##2.1 go build方式与go run方式区别?
1. go build 方式生成的可执行文件在没有go运行环境的机子上也可以执行,go run则依赖go运行环境
2. go build是可执行文件会包含go的第三方库,因此可执行文件会很大
## 2.2 go build生成其他名称的可执行文件
```
go build -o xxxx.exe  xx.go
```
这样就会生成xxx.exe文件
#3 注意
1. go应用程序以main()方法为入口
2. go方法由一条条语句构成,如果语句后不带分号";"则不能写在同一行
3. go定义的变量或者import的包没有使用,则编译无法通过
#4 gofmt使用
```
gofmt xx.go // 这样做只是将格式或后的内容输出

gofmt -w xx.go //格式化代码后写入文件
```
#5 '\r'的使用
```
fmt.Println("hello\rgo") //会输出gollo,因为'\r'后的会在第一个位置重新输出,并覆盖之前的
```

