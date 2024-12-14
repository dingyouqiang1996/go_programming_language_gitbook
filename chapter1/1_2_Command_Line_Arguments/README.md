# 1.2 命令行参数
- `os`包提供函数和值来处理系统命令行
  - 命令行参数由 `os.Args` 获得
  - `os.Args` 变量是字符串切片
  - 命令行参数是切片 `os.Args[1:]`
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
  - 注释写法为 `//`
  - `var` 声明字符串变量 `s` 和 `sep`
    - `var` 会部分初始化, 未赋值时数值为0或字符串为空
  - `+` 运算符 会将数值相加或字符串拼接
  - `:=` 是短变量声明, 他会基于初始化的值给予合适的类型(区别于`var`)
  - `i++` 意味值 `i`自增, 等价于 `i += 1` 
  - for循环的写法
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