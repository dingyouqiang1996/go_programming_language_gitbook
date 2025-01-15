# 2.3 Variables
- var 声明创建一个特定类型的变量，为其附加一个名称，并设置其初始值 `var name type = expression`
  - 类型或 = 表达式部分可以省略，但不能同时省略两者
  - 如果省略类型，则由初始化表达式确定
  - 如果省略表达式，则初始值是该类型的零值
    - 对于数字是 0
    - 对于布尔值是 false
    - 对于字符串是 ""
    - 对于接口和引用类型（切片、指针、map、通道、函数）是 nil
  - 像数组或结构体这样的聚合类型的零值是其所有元素或字段的零值
- 零值机制确保变量始终持有其类型的明确定义的值
- 在 Go 中，不存在未初始化的变量这样的东西
  - 这简化了代码，并且通常在无需额外工作的情况下确保边界条件的合理行为
```go
var s string
fmt.Println(s) // ""
```
- 可以在一个声明中声明并可选地初始化一组变量，使用匹配的表达式列表。省略类型允许声明不同类型的多个变量
```go
var i, j, k int // int, int, int
var b, f, s = true, 2.3, "four" // bool, flota64, string
```
- 初始化器可以是字面量值或任意表达式
- 包级变量在 main 开始之前初始化
- 一组变量也可以通过调用返回多个值的函数来初始化：
```go
var f, err = os.Open(name) // os.Open returns a file and an error
```
## 2.3.1 Short Variable Declarations
- var 声明通常用于需要显式类型与初始化表达式类型不同的局部变量
  - 或者当变量将在稍后被赋值且其初始值不重要时
```go
i := 100 // an int
var boiling float64 = 100 // a float64
var names []string
var err error
var p Point
```
- `:=` 可以声明并初始化多个变量
  - 带有多个初始化表达式的声明应该只在它们有助于可读性时使用
  - 例如用于短小且自然的组合，如 for 循环的初始化部分
```go
i, j := 0, 1
```
- 请记住，:= 是声明，而 = 是赋值
  - 多变量声明不应与元组赋值混淆
  - 在元组赋值中，左手边的每个变量都被赋予右手边的相应值
```go
i, j = j, i // swap values of i and j
```
- 像普通的 `var` 声明一样，短变量声明可用于调用像 `os.Open` 这样返回两个或更多值的函数
```go
f, err := os.Open(name)
if err != nil {
    return err
}
// ...use f...
f.Close()
```
- 短变量声明必须声明至少一个新变量，因此这段代码将无法编译
```go
f, err := os.Open(infile)
f, err := os.Create(outfile) // compile error: no new variables
```
## 2.3.2 Pointers
- 变量是包含值的存储单元，由声明创建的变量通过名称来识别
  - 如 x，但许多变量仅通过表达式来识别
  - 如 x[i] 或 x.f
- 指针值是变量的地址, 因此指针是值存储的位置
  - 并非每个值都有地址，但每个变量都有
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

## 2.3.3 The new Function
- 另一种创建变量的方法是使用内置函数 new。表达式 new(T) 创建一个未命名的 T 类型变量，将其初始化为 T 的零值，并返回其地址，该地址是 *T 类型的值。
```go
p := new(int) // p, of type *int, points to an unnamed int variable
fmt.Println(*p) // "0"
*p = 2 // sets the unnamed int to 2
fmt.Println(*p) // "2"
```
- 使用 new 创建的变量与取地址的普通局部变量没有区别，只是不需要发明（和声明）一个临时名称，我们可以在表达式中使用 new(T)。因此，new 只是一个语法上的便利，而不是一个基本概念：下面的两个 newInt 函数具有相同的行为
```go
func newInt() *int {
    return new(int)
}

func newInt() *int {
    var dummy int
    return &dummy
}
```
- 每次调用 new 都返回一个具有唯一地址的不同变量。
```go
p := new(int)
q := new(int)
fmt.Println(p == q) // "false"
```
- 有一个例外：两个类型不携带任何信息且因此大小为零的变量，如 struct{} 或 [0]int，可能根据实现具有相同的地址。
- new 函数相对较少使用，因为最常见的未命名变量是结构体类型，对于这些类型，结构体字面量语法更灵活
- 由于 new 是一个预声明的函数，一个关键字，因此可以在函数内将其名称重新定义为其他东西，例如：
```go
func delta(old, new int) int { return new - old }
```
- 当然，在 delta 函数内，内置的 new 函数不可用。

## 2.3.4 Lifetime of Variables
- 变量的生命周期是程序执行期间变量存在的那段时间间隔
- 包级变量的生命周期是整个程序的执行过程。相比之下，局部变量具有动态生命周期：每次执行声明语句时都会创建一个新的实例，并且变量一直存在，直到它变得不可达，此时其存储空间可以被回收。函数参数和结果也是局部变量；它们每次调用其包含的函数时都会被创建
- 例如 1.4节的Lissajous程序
  - 变量 t 每次 for 循环开始时都会被创建，并且每次循环迭代都会创建新的变量 x 和 y
```go
for t := 0.0; t < cycles*2*math.Pi; t += res {
    x := math.Sin(t)
    y := math.Sin(t*freq + phase)
    img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
}
```
- 垃圾回收器如何知道变量的存储空间可以被回收？完整的解释比我们这里需要的要详细得多，但基本思想是：每个包级变量和每个当前活跃函数的局部变量都可能是通往所讨论变量的路径的起点或根，沿着指针和其他类型的引用最终到达该变量。如果没有这样的路径存在，变量就变得不可达，因此它不再影响其余的计算
- 因为变量的生命周期仅由其是否可达来决定，局部变量可能会比其包含的循环的单次迭代存在得更久。即使其包含的函数已经返回，它也可能继续存在。
- 编译器可以选择在堆上或栈上分配局部变量，但或许令人惊讶的是，这个选择并不是由使用 var 或 new 来声明变量决定的
```go
var global *int
func f() {
    var x int
    x = 1
    global = &x
}

func g() {
    y := new(int)
    *y = 1
}
```
- 在这里，x 必须在堆上分配，因为它在 f 返回后仍然可以从变量 global 访问到，尽管它是作为局部变量声明的；我们说 x 从 f 中逃逸了
- 相反，当 g 返回时，变量 *y 变得不可达，可以被回收
- 由于 *y 没有从 g 中逃逸，即使它是用 new 分配的，编译器也可以安全地将 *y 分配在栈上
- 无论如何，逃逸的概念并不是你需要担心以编写正确代码的东西，尽管在性能优化期间最好记住这一点，因为每个逃逸的变量都需要额外的内存分配
- 垃圾回收在编写正确程序方面提供了极大的帮助，但它并没有解除你思考内存的负担。你不需要显式地分配和释放内存，但要编写高效的程序，你仍然需要意识到变量的生命周期。例如，将不必要的指向短命对象的指针保留在长寿命对象中，尤其是全局变量中，将阻止垃圾回收器回收这些短命对象

