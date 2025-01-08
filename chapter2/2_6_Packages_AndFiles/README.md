# 2.6 Packages and Files
- Go 中的包与其它语言中的库或模块具有相同的目的，支持模块化、封装、独立编译和重用。包的源代码位于一个或多个 .go 文件中，通常在以导入路径结尾的目录中；例如，gopl.io/ch1/helloworld 包的文件存储在目录 $GOPATH/src/gopl.io/ch1/helloworld 中
- 每个包为其声明提供一个独立的命名空间。例如，在 image 包中，标识符 Decode 指的是一个与 unicode/utf16 包中相同标识符不同的函数。要从包外引用函数，我们必须限定标识符以明确我们是指 image.Decode 还是 utf16.Decode
- 包还可以通过控制哪些名称在包外可见或导出来隐藏信息。在 Go 中，一个简单的规则决定了哪些标识符被导出，哪些没有：导出的标识符以大写字母开头。
- 为了说明基本概念，假设我们的温度转换软件变得流行，我们想将其作为新包提供给 Go 社区。我们该如何做
- 让我们创建一个名为 gopl.io/ch2/tempconv 的包，这是对前面示例的变体。（这里我们打破了通常按顺序编号示例的规则，以便包路径可以更现实。）该包本身存储在两个文件中，以展示如何访问包中不同文件中的声明；在现实生活中，像这样的小包只需要一个文件
- 我们把类型、常量和方法的声明放在了 tempconv.go 中
- 每个文件都以一个包声明开始，定义包名。当包被导入时，其成员被引用为 tempconv.CToF 等等。像类型和常量这样的包级名称在包的一个文件中声明，对包的其他文件可见，就好像源代码都在一个文件中一样。注意 tempconv.go 导入了 fmt，但 conv.go 没有，因为它没有使用 fmt 中的任何东西
- 因为包级 const 名称以大写字母开头，它们也可以通过限定名称如 tempconv.AbsoluteZeroC 访问
- 要在导入 gopl.io/ch2/tempconv 的包中将摄氏温度转换为华氏温度，我们可以编写以下代码：
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
- 每个文件都以一个包声明开始，定义包名。当包被导入时，其成员被引用为 tempconv.CToF 等等。像类型和常量这样的包级名称在包的一个文件中声明，对包的其他文件可见，就好像源代码都在一个文件中一样。注意 tempconv.go 导入了 fmt，但 conv.go 没有，因为它没有使用 fmt 中的任何东西
- 因为包级 const 名称以大写字母开头，它们也可以通过限定名称如 tempconv.AbsoluteZeroC 访问
- 紧接在包声明之前的文档注释记录了整个包。按照惯例，它应该以示例中所示的风格开始于一个概述句。每个包中只有一个文件应该有包文档注释。详细的文档注释通常放在一个单独的文件中，按照惯例称为 doc.go

## 2.6.1 Imports
- 在 Go 程序中，每个包都由一个称为导入路径的唯一字符串标识。这些字符串出现在像 "gopl.io/ch2/tempconv" 这样的导入声明中。语言规范没有定义这些字符串来自何处或它们的含义；由工具来解释它们。使用 go 工具（第 10 章）时，导入路径表示包含一个或多个 Go 源文件的目录，这些文件共同构成了包
- 除了导入路径外，每个包还有一个包名，这是出现在其包声明中的短名称（不一定是唯一的）。按照惯例，包名与其导入路径的最后一段匹配，这使得很容易预测 gopl.io/ch2/tempconv 的包名是 tempconv
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
- 导入声明将一个短名称绑定到导入的包，可以在整个文件中用来引用其内容。上面的导入让我们可以通过使用限定标识符如 tempconv.CToF 来引用 gopl.io/ch2/tempconv 中的名称。默认情况下，短名称是包名——在这个例子中是 tempconv——但导入声明可以指定一个替代名称以避免冲突
- 更好的做法是使用 golang.org/x/tools/cmd/goimports 工具，它会根据需要自动在导入声明中插入和删除包；大多数编辑器可以配置为每次保存文件时运行 goimports。像 gofmt 工具一样，它也会以规范格式美化 Go 源文件。



