# 2.6 Packages and Files
- Go 中的包与其它语言中的库或模块具有相同的目的
  - 支持模块化、封装、独立编译和重用
  - 包的源代码位于一个或多个 .go 文件中
  - 通常在以导入路径结尾的目录中
  - 例如，`gopl.io/ch1/helloworld` 包的文件存储在目录 `$GOPATH/src/gopl.io/ch1/helloworld` 中
- 每个包为其声明提供一个独立的命名空间
  - 例如，在 image 包中，标识符 Decode 指的是一个与 unicode/utf16 包中相同标识符不同的函数
  - 要从包外引用函数，我们必须限定标识符以明确我们是指 `image.Decode` 还是 `utf16.Decode`
- 包还可以通过控制哪些名称在包外可见或导出来隐藏信息
  - 在 Go 中，导出的标识符以大写字母开头
- 为了说明基本概念，假设我们的温度转换软件变得流行，我们想将其作为新包提供给 Go 社区。我们该如何做
- 让我们创建一个名为 gopl.io/ch2/tempconv 的包，这是对前面示例的变体
  - 这里我们打破了通常按顺序编号示例的规则，以便包路径可以更现实
  - 该包本身存储在两个文件中，以展示如何访问包中不同文件中的声明；在现实生活中，像这样的小包只需要一个文件
- 要在导入 `gopl.io/ch2/tempconv` 的包中将摄氏温度转换为华氏温度，我们可以编写以下代码：
- `gopl.io/ch2/tempconv/tempconv.go`
```go
// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
    AbsoluteZero Celsius = -273.15
    FreezingC Celsius = 0
    BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
```
- `gopl.io/ch2/tempconv/conv.go`
```go
package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```
- 每个文件都以一个包声明开始，定义包名
  - 当包被导入时，其成员被引用为 tempconv.CToF 等等
  - 像类型和常量的包级名称在包的一个文件中声明对包的其他文件可见，就像源代码都在一个文件一样
  - 注意 tempconv.go 导入了 fmt，但 conv.go 没有，因为它没有使用 fmt 中的任何东西
- 紧接在包声明之前的文档注释记录了整个包
  - 按照惯例，它应该以示例中所示的风格开始于一个概述句
  - 每个包中只有一个文件应该有包文档注释。详细的文档注释通常放在一个单独的文件中，按照惯例称为 `doc.go`

## 2.6.1 Imports
- 在 Go 程序中，每个包都由一个称为导入路径的唯一字符串标识
  - 这些字符串出现在像 "gopl.io/ch2/tempconv" 这样的导入声明中
  - 语言规范没有定义这些字符串来自何处或它们的含义
  - 由工具来解释它们。使用 go 工具（第 10 章）时，导入路径表示包含一个或多个 Go 源文件的目录，这些文件共同构成了包
- 除了导入路径外，每个包还有一个包名
  - 这是出现在其包声明中的短名称（不一定是唯一的）
  - 按照惯例，包名与其导入路径的最后一段匹配，这使得很容易预测 `gopl.io/ch2/tempconv` 的包名是 `tempconv`
- `gopl.io/ch2/cf/main.go`
```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"gopl.io/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
```
- 导入声明将一个短名称绑定到导入的包，可以在整个文件中用来引用其内容
  - 上面的导入让我们可以通过使用限定标识符如 tempconv.CToF 来引用 gopl.io/ch2/tempconv 中的名称
  - 默认情况下，短名称是包名——在这个例子中是 tempconv——但导入声明可以指定一个替代名称以避免冲突
- 更好的做法是使用 `golang.org/x/tools/cmd/goimports` 工具
  - 它会根据需要自动在导入声明中插入和删除包
  - 大多数编辑器可以配置为每次保存文件时运行 goimports
  - 像 gofmt 工具一样，它也会以规范格式美化 Go 源文件
## 2.6.2 Package Initialization
- 包初始化首先按声明顺序初始化包级变量，但依赖关系会先解决
```go
var a = b + c // a initialized third, to 3
var b = f() // b initialized second, to 2, by calling f
var c = 1 // c initialized first, to 1
func f() int { return c + 1 }
```
- 如果包有多个 .go 文件，它们将按照文件被提供给编译器的顺序进行初始化
  - go 工具在调用编译器之前会按名称对 .go 文件进行排序
  - 每个在包级声明的变量都以它的初始化表达式的值（如果有）开始其生命周期
- 但对于某些变量，如数据表，初始化表达式可能不是设置其初始值的最简单方式
  - 在这种情况下，init 函数机制可能更简单
  - 任何文件都可以包含任意数量的函数，其声明仅为
```go
func init() { /* ... */ }
```
- 这样的 init 函数不能被调用或引用，但除此之外它们是正常的函数
  - 在每个文件中，init 函数在程序启动时按声明顺序自动执行
  - 一次初始化一个包，顺序是程序中的导入顺序，先解决依赖关系
  - 导入q的包p可以确信q在p的初始化开始之前已完全初始化
  - 初始化从下到上进行
  - main 包是最后一个被初始化的
  - main 函数开始之前，所有包都已完全初始化
- 下面的包定义了一个名为 PopCount 的函数
  - 该函数返回 uint64 值中设置的位数，即值为 1 的位数
  - 这被称为其人口计数
  - 它使用 init 函数预先计算每个可能的 8 位值的结果表 pc
  - 以便 PopCount 函数不需要进行 64 步操作
  - 而只需返回八个表查找的总和
  - 这绝对不是计算位数的最快算法，但它方便于说明 init 函数，并展示了如何预先计算值表，这通常是很有用的编程技巧
- `gopl.io/ch2/popcount/popcount.go`
```go
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))]+
		pc[byte(x>>(2*8))]+
		pc[byte(x>>(3*8))]+
		pc[byte(x>>(4*8))]+
		pc[byte(x>>(5*8))]+
		pc[byte(x>>(6*8))]+
		pc[byte(x>>(7*8))])
}
```
- 注意 `for i := range pc` 这里i表示索引
