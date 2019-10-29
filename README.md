# 0 安装
可以访问<html src="https://golang.org">https://golang.org</html>下载安装
# 1 环境变量配置
windows下需要配置GOROOT(go语言安装目录),GOPATH(go程序目录)并在path中添加%GOROOT%/bin
暂时还不知道为什么添加GOPATH?
# 2 go执行
```
go build xx.go //编译go源代码 => 得到xx.exe文件
xx.exe //执行

go run xx.go //直接执行
```
## 2.1 go build方式与go run方式区别?
1. go build 方式生成的可执行文件在没有go运行环境的机子上也可以执行,go run则依赖go运行环境
2. go build是可执行文件会包含go的第三方库,因此可执行文件会很大
## 2.2 go build生成其他名称的可执行文件
```
go build -o xxxx.exe  xx.go
```
这样就会生成xxx.exe文件
# 3 注意
1. go应用程序以main()方法为入口
2. go方法由一条条语句构成,如果语句后不带分号";"则不能写在同一行
3. go定义的变量或者import的包没有使用,则编译无法通过
# 4 gofmt使用
```
gofmt xx.go // 这样做只是将格式或后的内容输出

gofmt -w xx.go //格式化代码后写入文件
```
# 5 '\r'的使用
```
fmt.Println("hello\rgo") //会输出gollo,因为'\r'后的会在第一个位置重新输出,并覆盖之前的
```

# 6 变量
Golang变量使用的三种方式:
1. 指定变量类型,声明后若不赋值,使用默认值(整数为0,字符串为空串)
    ```golang
    var i int //此时i为0
    ```
2. 根据值自行判定变量类型(类型推导)
    ```golang
    var num = 10.31
    ```
3. 省略var, :=左侧的变量不应该是已经声明过的,否则会导致编译错误
    ```golang
    num := "Tom"//等价于 var num String;num = "Tom";
    ```
4. Golang也支持一次声明多个变量
    ```golang
    var n1,n2,n3 int

    var n1,name,n3 = 100,"tom",888 // n1 => 100,name => "tom",n3 => 888

    n1,name,n3 := 100,"tom",888 // n1 => 100,name => "tom",n3 => 888
    ```
5. 声明全局变量
    ```golang
    package main
    
    var (
        i = 100
        n = 200
    )


    func main(){

    }
    
    //全局变量可以不使用
    ```
6. 该区域的数据值可以在同一类型下不断变化
    不能像js一样改变数据类型

7. 变量的数据类型
    数据类型
    1. 基本数据类型
    1).数值型:
        i. 整数类型 int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,byte
        ii. 浮点类型 float32,float64
    2)字符型,没有专门的字符型,使用byte保存字符
    3)bool型
    4)字符串
    2. 派生复杂数据类型
    1)指针(pointer)
    2)数组
    3)结构体(struct)
    4)管道(channel)
    5)函数
    6)切片(slice)
    7)接口(interface)
    8)map

        

# 7 程序中+号使用
1. 如果两边都是数值,做加法运算
2. 如果两边都是字符串,做字符串拼接
3. 一遍数值,一遍字符串会报错











