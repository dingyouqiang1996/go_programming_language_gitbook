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
## 2.3.1 Short Variable Declarations
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
## 2.3.2 Pointers
- 变量是包含值的存储单元，通过声明创建的变量由名称标识
- 指针值是变量的地址。因此，指针是值存储的位置。并非每个值都有地址，但每个变量都有
- 通过指针，我们可以间接地读取或更新变量的值，而无需使用甚至知道变量的名称，如果它确实有名称的话
- 如果声明了一个变量 var x int，表达式 &x（“x 的地址”）将返回一个指向整数变量的指针，即 *int 类型的值，读作“指向 int 的指针”
- 如果这个值称为 p，我们说“p 指向 x”或者等价地说“p 包含 x 的地址”
- p 指向的变量写作 *p。表达式 *p 返回该变量的值，一个 int，但由于 *p 表示一个变量，它也可以出现在赋值语句的左手边，在这种情况下，赋值语句更新该变量。
```go
x := 1
p := &x // p, of type *int, points to x
fmt.Println(*p) // "1"
*p = 2 // equivalent to x = 2
fmt.Println(x) // "2"
```
- 聚合类型变量的每个组成部分——结构体的一个字段或数组的一个元素——也是一个变量，因此也有一个地址
- 变量有时被描述为可寻址的值。表示变量的表达式是唯一可以应用取地址操作符 & 的表达式。
- 任何类型的指针的零值是 nil。如果 p 指向一个变量，则 p != nil 为真。指针是可比较的；两个指针相等当且仅当它们指向同一个变量或两者都是 nil
```go
var x, y int
fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"
```
- 函数返回局部变量的地址是完全安全的。例如，在下面的代码中，由对 f 的此次调用创建的局部变量 v 即使在调用返回后仍将存在，指针 p 仍将指向它：
```go
var p = f()
func f() *int {
    v := 1
    return &v
}
```
- 因为指针包含变量的地址，将指针参数传递给函数可以使函数更新间接传递的变量。例如，此函数将其参数指向的变量递增，并返回变量的新值，以便在表达式中使用。
```go
func incr(p *int) int {
    *p++ // increments what p points to; does not change p
    return *p
}
v := 1
incr(&v) // side effect: v is now 2
fmt.Println(incr(&v)) // "3" (and v is 3)
```
- 每次我们取变量的地址或复制指针时，我们都会创建新的别名或识别同一变量的方式。例如，*p 是 v 的别名。指针别名很有用，因为它允许我们不使用变量的名称来访问变量，但这是一把双刃剑：要找到访问变量的所有语句，我们必须知道它的所有别名。不仅仅是指针会创建别名；当我们复制其他引用类型的值，如切片、映射和通道时，也会发生别名，甚至包含这些类型的结构体、数组和接口也是如此。
- 指针是 flag 包的关键，该包使用程序的命令行参数来设置分布在程序中的某些变量的值。为了说明，这个对早期 echo 命令的变体接受两个可选标志：-n 使 echo 省略通常会打印的尾随换行符，-s sep 使其用字符串 sep 的内容而不是默认的单个空格分隔输出参数。由于这是我们的第四个版本，包名为 gopl.io/ch2/echo4
```go
// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
```
- 函数 flag.Bool 创建一个新的 bool 类型的标志变量。它接受三个参数：标志的名称（"n"）、变量的默认值（false）以及如果用户提供了无效参数、无效标志或 -h 或 -help 时将打印的消息。类似地，flag.String 接受一个名称、一个默认值和一个消息，并创建一个字符串变量。变量 sep 和 n 是指向标志变量的指针，必须通过 *sep 和 *n 间接访问。
- 

