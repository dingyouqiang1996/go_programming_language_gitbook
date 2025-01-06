# 2.3 Variables
- var 声明创建一个特定类型的变量，为其附加一个名称，并设置其初始值
  - 类型或 = 表达式部分可以省略，但不能同时省略两者
  - 如果省略类型，则由初始化表达式确定。如果省略表达式，则初始值是该类型的零值，对于数字是 0，对于布尔值是 false，对于字符串是 ""，对于接口和引用类型（切片、指针、映射、通道、函数）是 nil
  - 像数组或结构体这样的聚合类型的零值是其所有元素或字段的零值
- 零值机制确保变量始终持有其类型的明确定义的值；在 Go 中，没有未初始化的变量这样的东西
```go
var s string
fmt.Println(s) // ""
```
- 可以在一个声明中声明并可选地初始化一组变量，使用匹配的表达式列表。省略类型允许声明不同类型的多个变量
```go
var i, j, k int // int, int, int
var b, f, s = true, 2.3, "four" // bool, flota64, string
```
- 包级变量在 main 开始之前初始化
- 一组变量也可以通过调用返回多个值的函数来初始化：
```go
var f, err = os.Open(name) // os.Open returns a file and an error
```
### 2.3.1 Short Variable Declarations
- var 声明通常用于需要显式类型与初始化表达式类型不同的局部变量，或者当变量将在稍后被赋值且其初始值不重要时
```go
i := 100 // an int
var boiling float64 = 100 // a float64
var names []string
var err error
var p Point
```
- `:=` 可以声明并初始化多个变量
```go
i, j := 0, 1
```
- 请记住，:= 是声明，而 = 是赋值。多变量声明不应与元组赋值混淆，在元组赋值中，左手边的每个变量都被赋予右手边的相应值
- 短变量声明必须声明至少一个新变量，因此这段代码将无法编译。
```go
f, err := os.Open(infile)
f, err := os.Create(outfile) // compile error: no new variables
```
