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
- `gopl.io/ch2/tempconv.go`
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
- `gopl.io/ch2/conv.go`
```go
package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```
