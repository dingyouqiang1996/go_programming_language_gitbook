# 2.5 Type Declarations
- 变量或表达式的类型定义了它可能取值的特性
  - 例如它们的大小（位数或元素数量，可能）
  - 它们的内部表示方式，可以对它们执行的内在操作
  - 以及与它们关联的方法
- 在任何程序中，都有共享相同表示但表示非常不同概念的变量
  - int 可以用来表示循环索引、时间戳、文件描述符或月份
  - float64 可以表示以米每秒为单位的速度或几种温度标尺之一的温度
  - 字符串可以表示密码或颜色的名称
- 类型声明定义了一个新的命名类型，该类型与现有类型的底层类型相同
- 命名类型提供了一种方法，将底层类型的不同的、可能是不兼容的用途分开，以便它们不会被无意中混合使用
```go
type name underlying-type
```
- 类型声明通常出现在包级，命名类型在整个包中可见，并且如果名称被导出（以大写字母开头），它也可以从其他包中访问
- `tempconv/tempconv0.go`
```go
// Package temconv performs Celsius and Fahrenheit temperature computaitons
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```
- 此包定义了两种类型，Celsius 和 Fahrenheit，用于两种温度单位
  - 尽管两者都有相同的底层类型 float64
  - 但它们不是同一种类型，因此不能进行比较或在算术表达式中组合
  - 区分类型可以避免错误
    - 例如无意中将两种不同标尺的温度组合在一起
	- 需要显式类型转换，如 Celsius(t) 或 Fahrenheit(t)，才能从 float64 转换
  - Celsius(t) 和 Fahrenheit(t) 是转换，而不是函数调用
  - 它们不会以任何方式改变值或表示，但它们使意义的变化明确
  - 另一方面，函数 CToF 和 FToC 在两种标尺之间进行转换，它们确实返回不同的值
- 对于每种类型 T
  - 都有一个对应的转换操作 T(x)，将值 x 转换为类型 T
  - 如果两种类型具有相同的底层类型，或者两者都是指向相同底层类型变量的未命名指针类型，则允许从一种类型到另一种类型的转换
  - 这些转换改变类型但不改变值的表示。如果 x 可以赋给 T，则允许转换，但通常是多余的
- 数值类型之间以及字符串和某些切片类型之间的转换也是允许的
  - 这些转换可能会改变值的表示
  - 例如，将浮点数转换为整数会丢弃任何小数部分
  - 将字符串转换为 []byte 切片会分配字符串数据的副本
  - 无论如何，转换在运行时永远不会失败
- 命名类型的底层类型决定了其结构和表示，并且也决定了它支持的内在操作集
  - 这些操作与直接使用底层类型时相同
  - 这意味着算术运算符对 Celsius 和 Fahrenheit 的作用与对 float64 的作用相同
```go
fmt.Printf("%g\n", BoilingC-FreezingC)// "100" °C
boilingF := CToF(BoilingC)
fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
fmt.Printf("%g\n", boilingF-FreezingC) // compile erorr: type mismatch
```
- 像 == 和 < 这样的比较运算符也可以用来比较命名类型的值与相同类型的另一个值，或者与底层类型的值。但是，不同命名类型的两个值不能直接比较。
```go
var c Celsius
var f Fahrenheit
fmt.Println(c == 0) // "true"
fmt.Println(f >= 0) // "true"
fmt.Println(c == f) // "compile error: type mismatch"
fmt.Println(c == Celsius(f)) // "true"
```
- 最后一个例子,尽管名称如此,类型转换 Celsius(f) 并不改变其参数的值，只是改变其类型
- 如果命名类型有助于避免反复写出复杂的类型，则可以提供符号上的便利
  - 当底层类型像 float64 这样简单时，优势很小
  - 但对于复杂的类型，优势很大，我们将在讨论结构体时看到这一点
- 命名类型还可以为该类型的值定义新的行为
  - 这些行为以与类型相关联的一组函数的形式表达，称为类型的方法
- 下面的声明中，Celsius 参数 c 出现在函数名称之前，为 Celsius 类型关联了一个名为 String 的方法，该方法返回 c 的数值后跟 °C
```go
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```
- 许多类型声明了这种形式的 String 方法，因为它控制了类型值在通过 fmt 包以字符串形式打印时的显示方式