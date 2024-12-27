# 1.2 命令行参数
- `os`包提供函数和值来处理系统命令行
  - 命令行参数由 `os.Args` 获得
  - `os.Args` 变量是字符串切片
    - 切片是动态长度的序列, `s[i]` 可访问单个元素, `s[m:n]` 可以访问连续的序列， 通过 `len(s)` 得到元素的个数
    - `os.Args[0]` 是程序本身的名字
    - 命令行参数是切片 `os.Args[1:]`
    - `[1:len(os.Args)]` 如果省略 len(os.Args) 则写为 `os.Args[1:]`
## echo命令
- 示范unix命令 `echo` 的实现
```golang
package main

import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}
```
- 运行结果
```shell
$ go run ./echo1.go a b c 
a b c
```
- `echo` 代码中:
  -  `//` 开头的注释会被编译器忽略, 惯例是在package声明前用注释描述这个包
  - `var` 声明字符串变量 `s` 和 `sep`
    - `var` 会部分初始化, 隐式初始化时数值为0或字符串为空
  - `+` 运算符 会将数值相加或字符串拼接
    - `s += sep + os.Args[i]` 将sep的旧值和os.Args[i]拼接
    - += 是一个赋值运算符。每个算术和逻辑运算符，如 + 或 *，都有一个对应的赋值运算符
  - `:=` 是短变量声明, 他会基于初始化的值给予合适的类型(区别于`var`)
  - `i++` 意味值 `i`自增, 等价于 `i += 1` 
- 这种写法不通过每次循环都打印来实现, 而是通过循环将字符串拼接，然后再打印出来完整的变量
### for循环的写法
- go中的for写法，三个组成部分里不能用括号 `()`
```go
for 初始化变量; 条件; 后续操作 {
// 循环体
}
```
- for的死循环写法(go没有`while`)
```go
for {
// ...
}
```
- 使用 **另一种for循环写法** 的 `echo` 实现
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
```
- `range` 会循环切片 `os.args[1:]` 里元素的索引和值
  - 在 `_, arg := range os.Args[1:]` 的情况, _ 会被忽略
  - 在 `i, arg := range os.Args[1:]` 的情况, i 和 arg 必须被使用
### 声明变量 s 的方式
- 以下的声明变量 `s` 的方式都是等价的
```go
s := ""
var s string
var s = ""
var s string = ""
```
  - `s := ""` 只能用在函数内, 无法用于包等级的变量, 隐式声明类型
  - `var s string` 将 s 初始化为 `0值`, 显示声明类型
  - `var s = ""` 的声明方式可以同时用于多个值
  - `var s string = ""` 会显示声明类型和值
## 练习
- 修改 `echo` 程序, 添加打印 os.Args[0] 的代码
- 修改 `echo` 程序, 打印每个参数的索引和值
- 计算不同版本 `echo` 的运行时间